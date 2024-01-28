package leetcode

//买卖股票的最佳时机，https://leetcode.cn/leetbook/read/top-interview-questions-easy/xn8fsh/
func maxProfit(prices []int) int {
	//先用最差的办法算出结果，时间复杂度是 O（n2）
	//这种方式其实做了很多重复的计算
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

//双指针法，其实就是一个记录指针记录当前访问的过的最小值，当前值 - 已经访问过的最小值 O(n)
func maxProfitV2(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

//反而用动态规划算法把题目搞复杂了
