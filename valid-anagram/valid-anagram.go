package leetcode

// https://leetcode.cn/problems/valid-anagram/description/
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 字母异位词 是通过重新排列不同单词或短语的字母而形成的单词或短语，通常只使用所有原始字母一次。

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[rune]int)
	for _, item := range s {
		maps[item]++
	}
	for _, item := range t {
		_, ok := maps[item]
		if !ok {
			return false
		}
		if maps[item] == 0 {
			return false
		}
		maps[item]--
	}

	return true
}
