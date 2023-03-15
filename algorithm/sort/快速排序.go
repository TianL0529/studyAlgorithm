package sort

func quickSort(a []int, begin int, end int) {
	if begin >= end {
		return
	}

	// 进行切分
	pivot := partition(a, begin, end)
	// 对左半部分进行快排
	quickSort(a, begin, pivot-1)
	// 对右半部分进行快排
	quickSort(a, pivot+1, end)
}

func partition(a []int, begin int, end int) int {
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位
	for i < j {
		if a[i] > a[begin] {
			a[i], a[j] = a[j], a[i] // 交换
			j--
		} else {
			i++
		}
	}

	if a[i] >= a[begin] {
		i--
	}
	a[begin], a[i] = a[i], a[begin]

	return i
}
