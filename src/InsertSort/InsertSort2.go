package main

func main(){
	Show(insertSort2(ProduceArr()))
	Check(insertSort2(ProduceArr()))
}
/**
1.      初始时，a[0]自成1个有序区，无序区为a[1..n-1]。令i=1

2.      将a[i]并入当前的有序区a[0…i-1]中形成a[0…i]的有序区间。

3.      i++并重复第二步直到i==n-1。排序完成。
 */
func insertSort2(a[100]int)[100]int{
	var j int
	for i:=1;i<len(a);i++{
		for j=i-1;j>=0;j--{
			if a[j]<=a[i]{
				break
			}
		}

		if j!=i-1{
			temp:=a[i]
			var k int
			for k=i-1;k>j;k--{
				a[k+1]=a[k]
			}
			a[k+1]=temp
		}
	}
	return a
}


