package c

import (
	"strconv"
	"fmt"

	"math/rand"
)

func Show(a[100]int){
	for i := 0; i <=99; i++ {
		fmt.Print(strconv.Itoa(a[i])+" ")
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
func Check(a [100]int)  bool{
	flag:=a[0]>a[len(a)-1]
	if(flag){
		for i:=0;i<len(a)-1;i++{
			if a[i]<a[i+1]{
				fmt.Println(false)
				return false
			}
		}
	}else{
		for i:=0;i<len(a)-1;i++{
			if a[i]>a[i+1]{
				fmt.Println(false)
				return false
			}
		}
	}
	fmt.Println(true)
	return true
}

func A(){

}