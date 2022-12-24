package leetcode

// https://leetcode.cn/problems/regular-expression-matching/description/
// 正则表达式匹配
func isMatch(s string, p string) bool {
	//其实主要分为三种情况
	//1、p[j]是字母，2、p[j]是.,3、p[j]是*
	//p[j]是字母，那么s[i]必须和p[j]相等，否则返回false
	//p[j]是.，那么s[i]必须是字母，否则返回false
	m := len(s)
	n := len(p)

	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	//初始化状态
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//这块是核心逻辑，也是相关的状态转移
			if p[j-1] == '*' {
				//这一步的状态方程怎么想都想不明白
				if match(i, j-1) {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				//如果不是*号，可以按照这种方式匹配
				if match(i, j) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}

		}
	}
	return dp[m][n]
}
