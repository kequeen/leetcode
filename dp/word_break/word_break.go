package leetcode

//https://leetcode.cn/problems/word-break/?favorite=2cktkvj
//本质上还是动态规划 dp[i] = dp[j] && check(s[j:i])
//其实还有一点需要注意的是，想清楚一些初始状态
func wordBreak(s string, wordDict []string) bool {
	//用一个map去存储dict
	sLen := len(s)
	wordMap := make(map[string]bool)
	dp := make([]bool, sLen+1)
	for _, value := range wordDict {
		wordMap[value] = true
	}
	dp[0] = true
	for i := 1; i <= sLen; i++ {
		for j := 0; j < i; j++ {
			//如果是可以拆分的话，其实不需要再遍历了
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sLen]
}
