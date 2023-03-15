package sort

// 选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。
// 但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾。
func selectionSort(a []int64) []int64 {
	if len(a) <= 1 {
		return a
	}

	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		flag := false
		for j := i + 1; j < n; j++ {
			if a[min] <= a[j] {
				continue
			}
			min = j
			flag = true
		}
		if flag {
			a[i], a[min] = a[min], a[i]
		}
	}
	return a
}
