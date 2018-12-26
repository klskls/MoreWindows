package main

import (
	"strconv"
	"fmt"
	"math/rand"
)

func main() {
	Show(insertSort(ProduceArr()))
	Check(insertSort(ProduceArr()))
}
/**
这样的代码太长了，不够清晰。现在进行一下改写，
将搜索和数据后移这二个步骤合并。即每次a[i]先和前面一个数据a[i-1]比较，
如果a[i] > a[i-1]说明a[0…i]也是有序的，无须调整。否则就令j=i-1,temp=a[i]。
然后一边将数据a[j]向后移动一边向前搜索，当有数据a[j]<a[i]时停止并将temp放到a[j + 1]处。
---------------------
作者：MoreWindows
来源：CSDN
原文：https://blog.csdn.net/MoreWindows/article/details/6665714
版权声明：本文为博主原创文章，转载请附上博文链接！
 */
func insertSort(a [100] int)[100]int {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		k:=i
		for k > 0 && a[k-1] > temp {
			a[k]=a[k-1]
			k--
		}
		a[k]=temp
	}
	return a
}

func Show(a [100]int) {
	for i := 0; i <= 99; i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
	fmt.Println()

}

// 产生随机数字
func ProduceArr() [100] int {
	var arr [100]int
	for i := 99; i >= 0; i-- {
		arr[i] = rand.Intn(100)
	}
	return arr;
}

//检查数组是否有序
func Check(a [100]int) bool {
	flag := a[0] > a[len(a)-1]
	if (flag) {
		for i := 0; i < len(a)-1; i++ {
			if a[i] < a[i+1] {
				return false
			}
		}
	} else {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				return false
			}
		}
	}
	fmt.Println(true)
	return true
}
