package main

import (
	"math/rand"
	"fmt"
	"strconv"
)

func main() {
	Show(bubbleSort(ProduceArr()))
	fmt.Println(check(bubbleSort(ProduceArr())))


}

func bubbleSort(a [100] int) [100]int{
	for j := 0; j < 99; j++ {
		for i := 0; i < 99-j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}

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
func check(a [100]int)  bool{
	flag:=a[0]>a[len(a)-1]
	if(flag){
		for i:=0;i<len(a)-1;i++{
			if a[i]<a[i+1]{
				return false
			}
		}
	}else{
		for i:=0;i<len(a)-1;i++{
			if a[i]>a[i+1]{
				return false
			}
		}
	}
	return true
}


