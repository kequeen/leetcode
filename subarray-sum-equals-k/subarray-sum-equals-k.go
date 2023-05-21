package leetcode

// https://leetcode.cn/problems/subarray-sum-equals-k/?envType=study-plan-v2&id=top-100-liked
// 和为k的子数组
// 我理解应该就是使用左右指针去处理这个问题
// 如果k不能是负值，其实滑动窗口应该是OK的
func subarraySum(nums []int, k int) int {
	left := 0
	right := 0
	sum := 0
	count := 0

	for left <= right && right < len(nums) {
		sum += nums[right]
		right++
		for sum > k {
			sum -= nums[left]
			left++
		}
		if sum == k && left != right {
			count++
		}
	}
	return count
}

// 暴力解法，遍历
func subarraySumV2(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 所谓采用前缀和的方式,还是稍微有点想不通
func subarraySumV3(nums []int, k int) int {
	count := 0
	preSum := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			count += m[preSum-k]
		}
		m[preSum]++
	}
	return count
}
