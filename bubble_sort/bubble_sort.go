package leetcode

// 冒泡排序
// 优化版本的冒泡排序，如果本轮没有发生交换，则提前退出
func bubble_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swap := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
