package leetcode

//硬币找零 有限硬币 判断是否能找零
func chargeLimit(money int, coins []int) bool {
	dp := make([]int, money+1)
	//初始化
	for i := 1; i <= money; i++ {
		dp[i] = money + 1
	}
	for i := range coins {
		for j := money; j >= coins[i]; j-- {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	return dp[money] != money+1
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
