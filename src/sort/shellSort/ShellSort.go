package shellSort

/**
希尔排序的实质就是分组插入排序，该方法又称缩小增量排序，
因DL．Shell于1959年提出而得名。

 

该方法的基本思想是：先将整个待排元素序列分割成若干个子序列
（由相隔某个“增量”的元素组成的）分别进行直接插入排序，然后依次缩减增量再进行排序，
待整个序列中的元素基本有序（增量足够小）时，再对全体元素进行一次直接插入排序。
因为直接插入排序在元素基本有序的情况下（接近最好情况），效率是很高的，
因此希尔排序在时间效率上比前两种方法有较大提高。
---------------------
作者：MoreWindows
来源：CSDN
原文：https://blog.csdn.net/morewindows/article/details/6668714
版权声明：本文为博主原创文章，转载请附上博文链接！
 */

func ShellSort1(a []int) [] int {
	for gap := len(a) / 2; gap > 0; gap = gap / 2 {
		for j := 0; j < gap; j++ {
			for k := j + gap; k < len(a); k = k + gap {
				temp := a[k]
				m := k - gap //m=j  注意这里m=j是不对的
				for ; m >= 0 && a[m] > temp; m = m - gap {
					a[m+gap] = a[m]
				}
				a[m+gap] = temp
			}
		}
	}
	return a
}

/**
从第二个元素直接开始插入排序

很明显，上面的shellsort1代码虽然对直观的理解希尔排序有帮助，
但代码量太大了，不够简洁清晰。因此进行下改进和优化，以第二次排序为例，
原来是每次从1A到1E，从2A到2E，可以改成从1B开始，先和1A比较，然后取2B与2A比较，
再取1C与前面自己组内的数据比较…….。这种每次从数组第gap个元素开始，
每个元素与自己组内的数据进行直接插入排序显然也是正确的。

 */

func ShellSort2(a []int) [] int {
	for gap := len(a) / 2; gap > 0; gap = gap / 2 {
		for k := gap; k < len(a); k = k + gap {
			temp := k
			i := k - gap
			for ; i >= 0 && temp < a[i]; i = i - gap {
				a[i+gap]=a[i]
			}
			a[i+gap]=temp
		}
	}
	return a
}
/**
再将直接插入排序部分用
白话经典算法系列之二 直接插入排序的三种实现
中直接插入排序的第三种方法来改写下
 */

func ShellSort3(a []int) [] int {
	for gap := len(a) / 2; gap > 0; gap = gap / 2 {
		for k := gap; k < len(a); k = k + gap {
			i := k - gap
			for ; i >= 0 && a[i+gap] < a[i]; i = i - gap {
				a[i+gap],a[i]=a[i],a[i+gap]
			}

		}
	}
	return a
}


