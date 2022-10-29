package leetcode

//https://leetcode.cn/problems/letter-combinations-of-a-phone-number/?favorite=2cktkvj
//电话号码的字母组合
func letterCombinations(digits string) []string {
	//这种组合问题就是标准的回溯法
	res := make([]string, 0)
	digitsMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pgrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var comb string
	var dfs func(string, int)
	dfs = func(s string, idx int) {
		//判定终止条件
		if len(comb) == len(digits) {
			res = append(res, string(comb))
			return
		}
		if idx == len(digits) {
			return
		}
		//下一个
		dfs(s, idx+1)

		//遍历当前
		current, ok := digitsMap[s[idx]]
		if !ok {
			return
		}
		for _, value := range current {
			comb = comb + string(value)
			dfs(s, idx+1)
			comb = comb[:len(comb)-1]
		}

	}
	dfs(digits, 0)

	return res
}
