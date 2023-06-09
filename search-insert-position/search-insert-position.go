package leetcode

// https://leetcode.cn/problems/search-insert-position/description/
// 搜索插入位置
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
