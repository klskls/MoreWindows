package main

import "fmt"

func main() {
	Show(bubbleSort3(ProduceArr()))
	fmt.Println(check(bubbleSort3(ProduceArr())))
}

/**
再做进一步的优化。如果有100个数的数组，仅前面10个无序，
后面90个都已排好序且都大于前面10个数字，
那么在第一趟遍历后，最后发生交换的位置必定小于10，
且这个位置之后的数据必定已经有序了，记录下这位置，
第二次只要从数组头部遍历到这个位置就可以了。

记录每次遍历的最后一次交换位置
 */
func bubbleSort3(a [100]int) [100]int {
	var temp int
	temp = 99
	for temp != 0 {
		k:=temp
		for i := 0; i < k /*99-j*/ ; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				temp = i
			}
			temp = temp - 1

		}
	}

	return a
}
