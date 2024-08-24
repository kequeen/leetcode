package leetcode

// https://leetcode.cn/problems/edit-distance/solutions/
// 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

// 你可以对一个单词进行如下三种操作：

// 插入一个字符
// 删除一个字符
// 替换一个字符
// 编辑距离，确实是特别经典的题目
// 最主要是想明白这个动态规划的公式
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	//如果有一个字符串为空的话
	if m*n == 0 {
		return m + n
	}
	//用动态规划的方式解决
	//本质上变更只有三种 A插入1个字符，B插入一个字符，A改变一个字符
	//因为A插入一个字符，其实等价于B删除一个字符
	//定义二维数组
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	//初始化
	dp[0][0] = 0
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果相等
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insert := dp[i][j-1] + 1    // 插入
				delete := dp[i-1][j] + 1    // 删除
				replace := dp[i-1][j-1] + 1 // 替换
				dp[i][j] = min(min(insert, delete), replace)
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
