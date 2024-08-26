package leetcode

// https://leetcode.cn/problems/longest-palindromic-substring/
// 最长回文子字符串
// 暴力方法肯定是超时的，
// 判断最长回文字符串，从当前出发，往左或者往右，其应该都相等
func longestPalindrome(s string) string {
	//我是感觉中心扩散的方式比较容易理解
	length := len(s)
	if length < 2 {
		return s
	}
	start := 0
	end := 0

	maxLen := 1
	for i := 0; i < length; i++ {
		//因为回文有两种情况
		len1 := calPalindrome(s, i, i, length)
		len2 := calPalindrome(s, i, i+1, length)
		currentLen := max(len1, len2)
		if currentLen > maxLen {
			maxLen = currentLen
			start = i - (currentLen-1)/2
			end = i + currentLen/2
		}
	}
	return s[start : end+1]
}

// 计算回文字符串长度
func calPalindrome(s string, left int, right int, length int) int {
	for left >= 0 && right < length && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
