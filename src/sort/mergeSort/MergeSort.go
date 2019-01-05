package mergeSort

/**
合并两个有序数组
 */
func MergeArray(a []int, b []int) [] int {
	c := make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < a[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = a[j]
			j++
		}
		k++
	}

	for i < len(a) {
		c[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		c[k] = a[j]
		k++
		j++
	}
	return c
}
