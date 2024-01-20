package leetcode

//存在重复元素2
//https://leetcode.cn/problems/contains-duplicate-ii/

// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：nums = [1,2,3,1], k = 3
// 输出：true
// 示例 2：

// 输入：nums = [1,0,1,1], k = 1
// 输出：true
// 示例 3：

// 输入：nums = [1,2,3,1,2,3], k = 2
// 输出：false

// 用一下滑动窗口的办法
// 可以删除超过位置的窗口，也可以更新上次出现的位置
func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	duplicateMap := make(map[int]bool)
	for i := 0; i < length; i++ {
		if i > k {
			delete(duplicateMap, nums[i-k-1])
		}
		if duplicateMap[nums[i]] {
			return true
		} else {
			duplicateMap[nums[i]] = true
		}
	}
	return false
}
