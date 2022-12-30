package leetcode

// 回文子串的数目
// https://leetcode.cn/problems/palindromic-substrings/?favorite=2cktkvj
// 我反而觉得dp的算法比他们各种算法都清晰明了得多
func countSubstrings(s string) int {
	//我理解这种也是一种动归，其实就是
	// if s[i] == s[j] && dp[i+1][j-1] == true ,则dp[i][j] == true
	count := 0
	sLen := len(s)
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
		dp[i][i] = true
		count++
	}
	for i := 1; i < sLen; i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				if i-j == 1 || i-j == 2 || dp[i-1][j+1] {
					dp[i][j] = true
					count++
				}
			}
		}
	}
	return count
}
