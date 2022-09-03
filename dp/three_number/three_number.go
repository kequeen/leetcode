package leetcode


//有三个数1，3，5，判断一个数有多少种组合方法
func three_number(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	
	
	return dp[n]
	
}
