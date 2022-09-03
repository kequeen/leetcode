package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnq4km/
// dp[n] = max(dp[n-1], dp[n-2] + nums[n])
func rob(nums []int) int {
	maxRob := 0
	numsLen := len(nums)
	if numsLen == 0 {
		return maxRob
	}
	if numsLen == 1 {
		return nums[0]
	}

	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxRob = max(nums[0], nums[1])

	dp := make([]int, numsLen)
	//初始化
	dp[0] = nums[0]
	dp[1] = maxRob

	for i := 2; i < numsLen; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		maxRob = max(dp[i], maxRob)
	}

	return maxRob
}
