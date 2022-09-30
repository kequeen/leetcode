package leetcode

import "strings"

//https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/
//至少有 K 个重复字符的最长子串

//需要将问题拆解成更小的问题，分解
//可以先统计哪些字符出现的频率小于k，答案肯定不包括这些字符串
func longestSubstring(s string, k int) int {
	maxLen := 0
	sLen := len(s)
	frequecyMap := make(map[byte]int)
	for i := 0; i < sLen; i++ {
		frequecyMap[s[i]]++
	}
	var splitChar byte
	endFlag := true
	for key, value := range frequecyMap {
		if value < k {
			splitChar = key
			endFlag = false
		}
	}

	//如果没有小于k的字符，那说明已经是符合条件的字符串
	if endFlag {
		return sLen
	}

	subStrs := strings.Split(s, string(splitChar))
	for _, value := range subStrs {
		maxLen = max(maxLen, longestSubstring(value, k))
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
