package leetcode

//https://leetcode.cn/problems/generate-parentheses/description/?favorite=2cktkvj
//括号生成，典型的回溯法题目

func generateParenthesis(n int) []string {
	var s string
	ans := make([]string, 0)
	var backtrace func(string, int, int, int)
	backtrace = func(s string, open int, close int, n int) {
		if len(s) == n*2 {
			ans = append(ans, s)
			return
		}
		if open < n {
			s = s + "("
			backtrace(s, open+1, close, n)
			s = s[:len(s)-1]
		}
		if open > close {
			s = s + ")"
			backtrace(s, open, close+1, n)
		}
	}

	backtrace(s, 0, 0, n)
	return ans
}
