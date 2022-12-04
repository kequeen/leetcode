package leetcode

// https://leetcode.cn/problems/water-bottles/
// 换水问题
func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + numEmptyWaterBottles(numBottles, numExchange)
}

// 空瓶换水的子问题
func numEmptyWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return 0
	}
	num := numBottles / numExchange
	left := numBottles % numExchange
	return numEmptyWaterBottles(num+left, numExchange) + num
}
