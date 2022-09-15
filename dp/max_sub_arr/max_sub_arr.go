package leetcode

//最大子序和 https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn3cg3/
// dp [i] = max(num[i], dp[i-1] + nums[i])
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	dp := make([]int, numsLen)
	maxSum := nums[0]
	dp[0] = nums[0]
	for i := 1; i < numsLen; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}
