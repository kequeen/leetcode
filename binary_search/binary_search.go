package leetcode

// 比较标准的二分查找
// https://leetcode.cn/problems/binary-search/
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
