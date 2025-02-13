package leetcode

// 硬币找零，返回需要的最少硬币数，无限找零硬币
func charge(money int, coins []int) int {
	dp := make([]int, money+1)
	//初始化
	for key, _ := range dp {
		dp[key] = money + 1
	}

	// dp 公式 dp(n) = min(dp[n], dp[n - value]) value为硬币的值
	for i := 1; i <= money; i++ {
		for _, value := range coins {
			if i-value > 0 {
				dp[i] = min(dp[i], dp[i-value]+1)
			} else if i-value == 0 {
				dp[i] = 1
			}
		}
	}
	return dp[money]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
