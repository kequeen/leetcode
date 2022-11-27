package leetcode

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/
//无重复的最长子串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//这个是有问题的版本
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	sLen := len(s)
	tempLen := 0
	strMap := make(map[byte]bool)
	for i := 0; i < sLen; i++ {
		if strMap[s[i]] {
			tempLen = 0
			strMap = map[byte]bool{}
		} else {
			tempLen++
			strMap[s[i]] = true
		}
		maxLen = max(maxLen, tempLen)
	}
	return maxLen
}

//其实也还是双指针，这种其实也是滑动窗口
func lengthOfLongestSubstringV2(s string) int {
	maxLen := 0
	sLen := len(s)
	//存储出现的次数
	sMap := make(map[byte]int)
	//滑动窗口，记录往右的最大偏移位置
	right := -1
	for i := 0; i < sLen; i++ {
		//往右移动
		if i > 0 {
			delete(sMap, s[i-1])
		}
		//其实就是一直往右，如果有重复的，本身就会停下来，会等左边的指针过来，删除掉重复的元素
		for right+1 < sLen && sMap[s[right+1]] == 0 {
			sMap[s[right+1]]++
			right++
		}
		maxLen = max(maxLen, right-i+1)
	}
	return maxLen
}

//再升级，不从-1开始了
func lengthOfLongestSubstringV3(s string) int {
	maxLen := 0
	sLen := len(s)
	//存储出现的次数
	sMap := make(map[byte]int)
	//滑动窗口，记录往右的最大偏移位置
	right := 0
	for i := 0; i < sLen; i++ {
		//往右移动
		if i > 0 {
			delete(sMap, s[i-1])
		}
		for right < sLen && sMap[s[right]] == 0 {
			sMap[s[right]]++
			right++
		}
		maxLen = max(maxLen, right-i)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
