package selectSort
/**
直接选择排序
 */
func SelectSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		minI := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minI] {
				minI = j
			}
		}
		if i != minI {
			a[i], a[minI] = a[minI], a[i]
		}
	}
	return a
}
