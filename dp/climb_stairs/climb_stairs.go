package leetcode

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn854d/
// 最经典的爬楼梯问题
func climbStairs(n int) int {
	//这种直接递归的方式会超时
	//dp公式 dp[n] = dp[n-1] + dp[n-2]
	//dp[1] = 1, dp[2] =2
	if n == 1 || n == 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 不重复计算，以空间换时间
func climbStairsV2(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1)
	//初始化
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 最佳解法
func climbStairsV3(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	//其实只需要存储三个变量，结果ans，前面的两个变量p,q
	p := 1
	q := 2
	ans := 0
	for i := 3; i <= n; i++ {
		ans = p + q
		p = q
		q = ans
	}
	return ans
}
