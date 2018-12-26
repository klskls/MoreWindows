package main

func main() {
	Show(insertSort3(ProduceArr()))
	Check(insertSort3(ProduceArr()))
}
/**
再对将a[j]插入到前面a[0…j-1]的有序区间所用的方法进行改写，
用数据交换代替数据后移。如果a[j]前一个数据a[j-1] > a[j]，
就交换a[j]和a[j-1]，再j--直到a[j-1] <= a[j]。
这样也可以实现将一个新数据新并入到有序区间。
 */
func insertSort3(a [100]int) [100]int {

	for i:=1;i<len(a);i++{

		for j:=i-1;j>0&&a[j]>a[i];j--{
			a[j],a[i]=a[i],a[j]
		}
	}
	return a
}
