package main

import "fmt"

func main(){
	Show(bubbleSort2(ProduceArr()))
	fmt.Println(check(bubbleSort2(ProduceArr())))
}
/*
下面对其进行优化，设置一个标志，如果这一趟发生了交换，
则为true，否则为false。明显如果有一趟没有发生交换，说明排序已经完成。
 */
func bubbleSort2(a [100]int )[100]int{

	for j:=0;j<99;j++{
		b:=true
		for i:=0;i<99-j;i++{
			b=false
			if a[i]>a[i+1]{
				a[i],a[i+1]=a[i+1],a[i]
				b=true
			}
		}
		if !b {
			break
		}
	}
	return a
}
