package mergeSort

/**
合并两个有序数组
https://blog.csdn.net/MoreWindows/article/details/6678165
 */
func MergeArray(a []int, b []int, c []int) [] int {
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < a[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, a[j])
			j++
		}
	}

	for i < len(a) {
		c = append(c, a[i])
		i++
	}
	for j < len(b) {
		c = append(c, a[j])
		j++
	}
	return c
}

func mergeArray(a []int, first int, mid int, last int, temp [] int) {
	i, j := first, mid+1
	k := 0
	for i <= mid && j <= last {
		if a[i] < a[j] {
			temp[k] = a[i]
			k++
			i++
		} else {
			temp[k] = a[j]
			k++
			j++
		}
		for i <= mid {
			temp[k] = a[i]
			k++
			i++
		}
		for j <= last {
			temp[k] = a[j]
			j++
			k++
		}
		for n := 0; n < k; n++ {
		 	a[first+n]=temp[n]
		}
	}
}

func MergeSort(a []int,first int,last int,temp [] int)  {
	if first<last{
		mid:=(first+last)/2
		MergeSort(a,first,mid,temp)
		MergeSort(a,mid+1,last,temp)
		mergeArray(a,first,mid,last,temp)
	}
}
