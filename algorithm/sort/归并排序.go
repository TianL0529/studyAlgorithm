package sort

// 如果要排序一个数组，我们先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了。
// 分治思想
// 写递归代码的技巧：分析得出递推公式，然后找到终止条件，最后将递推公式翻译成递归代码
func mergeSort(a []int64) []int64 {
	if len(a) <= 1 {
		return a
	}

	// 获取分区位置
	p := len(a) / 2
	// 通过递归分区
	left := mergeSort(a[0:p])
	right := mergeSort(a[p:])
	// 排序后合并
	return merge(left, right)
}

func merge(left []int64, right []int64) []int64 {
	i, j := 0, 0
	m, n := len(left), len(right)

	var result []int64
	for {
		// 任何一个区间遍历完，则退出
		if i >= m || j >= n {
			break
		}
		// 对所有区间数据进行排序
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 如果左侧区间还没有遍历完，将剩余数据放到结果集
	if i != m {
		for ; i < m; i++ {
			result = append(result, left[i])
		}
	}

	// 如果右侧区间还没有遍历完，将剩余数据放到结果集
	if j != n {
		for ; j < n; j++ {
			result = append(result, right[j])
		}
	}

	// 返回排序后的结果集
	return result
}
