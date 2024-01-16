package leetcode

//加油站
//https://leetcode.cn/leetbook/read/top-interview-questions/xmguej/

// 被chatgpt教写题了,不过确实比原来的好太多了，原来的我还考虑一圈的情况，实际上不需要考虑的
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalTank, currTank, start := 0, 0, 0

	for i := 0; i < n; i++ {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]

		if currTank < 0 {
			// 无法从当前加油站出发，重置起始位置并清空当前累计油箱量
			start = i + 1
			currTank = 0
		}
	}
	if totalTank >= 0 {
		// 环路可以绕行一周
		return start
	}
	return -1
}
