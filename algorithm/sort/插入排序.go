package sort

// 首先，我们将数组中的数据分为两个区间，已排序区间和未排序区间。初始已排序区间只有一个元素，就是数组的第一个元素。
// 插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。
// 重复这个过程，直到未排序区间中元素为空，算法结束。
func insertionSort(a []int64) []int64 {
	if len(a) <= 1 {
		return a
	}

	n := len(a)
	for i := 0; i < n-1; i++ {
		value := a[i+1]
		j := i
		for ; j >= 0; j-- {
			if a[j] <= value {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = value
	}
	return a
}
