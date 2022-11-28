package leetcode

//https://leetcode.cn/problems/longest-common-prefix/
//最长公共前缀
func longestCommonPrefix(strs []string) string {
	//其实最简单的方法就是直接遍历
	var ans string
	if len(strs) == 0 {
		return ans
	}
	ans = strs[0]
	for i := 1; i < len(strs); i++ {
		ans = getCommonPrefix(ans, strs[i])
		if len(ans) == 0 {
			break
		}
	}

	return ans
}

func getCommonPrefix(str1, str2 string) string {
	var ans string
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] == str2[i] {
			ans += string(str1[i])
		} else {
			break
		}
	}
	return ans
}
