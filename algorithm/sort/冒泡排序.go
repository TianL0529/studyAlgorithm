package sort

// 冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。
// 如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作。
func bubbleSort(a []int64) []int64 {
	if len(a) <= 1 {
		return a
	}

	n := len(a)
	for j := 0; j < n; j++ {
		flag := false
		for i := 0; i < n-j-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	return a
}
