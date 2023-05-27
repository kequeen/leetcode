package leetcode

// 最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/?envType=study-plan-v2&envId=top-100-liked
func longestCommonSubsequence(text1 string, text2 string) int {
	maxLen := 0
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1)
	for i, _ := range dp {
		dp[i] = make([]int, len2)
	}
	//初始化
	for i := 0; i < len1; i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}
	for j := 0; j < len2; j++ {
		if text2[j] == text1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	//动态规划公式 dp[i][j] = max(dp[i-1][j], dp[i][j-1], dp[i-1] + dp[j-1]) + cur
	// if text[i] == text2[j], then cur = 1 else cur = 0
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = max(max(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			} else {
				dp[i][j] = max(max(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 上面自己的写法建模有些复杂了，其实可以简化一下，因为dp[i][j]只和dp[i-1][j], dp[i][j-1]
// 为什么不需要比较dp[i-1][j-1]，因为dp[i-1][j] 和 dp[i][j-1]都和这个有关
func longestCommonSubsequenceV2(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len1][len2]
}
