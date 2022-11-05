package leetcode

//https://leetcode.cn/problems/jump-game/?favorite=2cktkvj
//跳跃游戏
//我理解其实就是一道动态规划的题目
//这个的时间复杂度是O(n^2),空间复杂度是O(n)
//直觉上应该有更佳的解法，看答案果然也有更佳的解法
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && j+nums[j]-i >= 0 {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

//更优的解法，其实就是维护一个当前能到达的最大的位置
//这个的时间复杂度是O(n),空间复杂度是O(1)
func canJumpV2(nums []int) bool {
	numsLen := len(nums)
	maxJump := 0
	for i := 0; i < numsLen; i++ {
		if i <= maxJump {
			maxJump = max(maxJump, i+nums[i])
			if maxJump >= numsLen-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
