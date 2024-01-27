package leetcode

import (
	"sort"
)

// https://leetcode.cn/problems/group-anagrams/?favorite=2cktkvj
// 字母异或类分词
// 其实我理解是需要一个map去将这个进行分类
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	strsMap := make(map[string][]string, 0)
	for _, str := range strs {
		s := []byte(str)
		//上面这种写法不行
		//sort.Slice(s, func(i, j int) bool { return s[i]-s[j] < 0 })
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		strsMap[sortedStr] = append(strsMap[sortedStr], str)
	}
	for _, v := range strsMap {
		res = append(res, v)
	}
	return res
}
