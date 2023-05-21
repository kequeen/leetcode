package leetcode

// 找到字符串中所有字母异位词
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
func findAnagrams(s string, p string) []int {
	var ans []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}
	sCount, pCount := [26]int{}, [26]int{}
	for i, v := range p {
		pCount[v-'a']++
		sCount[s[i]-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	//采用滑动窗口进行比较
	for i := 0; i < sLen-pLen; i++ {
		//第一位退场
		sCount[s[i]-'a']--
		//末位进场
		sCount[s[i+pLen]-'a']++
		if pCount == sCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}
