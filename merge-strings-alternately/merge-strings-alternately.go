package leetcode

// https://leetcode.cn/problems/merge-strings-alternately/?envType=study-plan-v2&id=leetcode-75
// 其实就是归并排序，只是在golang中字符串的处理比较麻烦
func mergeAlternately(word1 string, word2 string) string {
	var result string
	var i, j int
	for i < len(word1) && j < len(word2) {
		result += string(word1[i])
		result += string(word2[j])
		i++
		j++
	}
	if i < len(word1) {
		result += word1[i:]
	}
	if j < len(word2) {
		result += word2[j:]
	}
	return result
}
