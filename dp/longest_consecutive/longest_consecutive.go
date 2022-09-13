package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions/x2xmre/
// 最长连续序列
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//使用map存储数字和对应的长度
	numsMap := make(map[int]int)
	dp := make(map[int]int)
	for _, value := range nums {
		numsMap[value] = 1
	}
	maxLen := 0
	for key := range numsMap {
		if dp[key-1] != 0 {
			dp[key] = dp[key-1] + 1
		} else {
			//有可能没有遍历完
			currentLen := 1
			currentKey := key
			for {
				if numsMap[currentKey-1] == 1 {
					currentLen++
					currentKey--
				} else {
					break
				}
			}
			//可以从key到currentKey全部算出来
			//其实核心还是要避免重复的计算,有重复的计算肯定会超时
			for i := key; i >= currentKey; i-- {
				dp[i] = currentLen
				currentLen--
			}

		}
		maxLen = max(maxLen, dp[key])
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
