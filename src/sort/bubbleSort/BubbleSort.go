package bubbleSort
//https://blog.csdn.net/morewindows/article/details/6657829

func BubbleSort(a [100] int) [100]int {
	for j := 0; j < 99; j++ {
		for i := 0; i < 99-j; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	return a
}

/*
下面对其进行优化，设置一个标志，如果这一趟发生了交换，
则为true，否则为false。明显如果有一趟没有发生交换，说明排序已经完成。
 */
func BubbleSort2(a [100]int )[100]int{
	var b bool=true
	for j:=0;b||j<99;j++{
		b=false
		for i:=0;i<99-j;i++{

			if a[i]>a[i+1]{
				a[i],a[i+1]=a[i+1],a[i]
				b=true
			}
		}

	}
	return a
}

/**
再做进一步的优化。如果有100个数的数组，仅前面10个无序，
后面90个都已排好序且都大于前面10个数字，
那么在第一趟遍历后，最后发生交换的位置必定小于10，
且这个位置之后的数据必定已经有序了，记录下这位置，
第二次只要从数组头部遍历到这个位置就可以了。

记录每次遍历的最后一次交换位置
 */
func BubbleSort3(a [100]int) [100]int {
	var temp int
	temp = 99
	for temp != 0 {
		k:=temp
		temp=0//没有发生交换
		for i := 0; i < k /*99-j*/ ; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				temp = i
			}
		}
	}
	return a
}
