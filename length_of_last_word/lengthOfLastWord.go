package leetcode

//https://leetcode.cn/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	length := 0
	//其实就是从后往前遍历，遍历到第一个不是空字符串的内容
	sLen := len(s)
	for i := sLen - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		begin := i
		for begin >= 0 && s[begin] != ' ' {
			length++
			begin--
		}
		break
	}
	return length
}
