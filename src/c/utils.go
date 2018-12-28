package c

import (
	"strconv"
	"fmt"

	"math/rand"
)

func Show(a[]int){
	for i := 0; i <=len(a)-1; i++ {
		fmt.Print(strconv.Itoa(a[i])+" ")
	}
	fmt.Println()

}
// 产生随机数字
func ProduceArr(n int) [] int {
	s:=make([]int,n)
	for i := 0; i <=n-1; i++ {
		s[i] = rand.Intn(n)
	}
	return s;
}

//检查数组是否有序
func Check(a []int)  bool{
	n:=len(a)
	flag:=a[0]>a[n-1]
	if(flag){
		for i:=0;i<n-1;i++{
			if a[i]<a[i+1]{
				fmt.Println(false)
				return false
			}
		}
	}else{
		for i:=0;i<n-1;i++{
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