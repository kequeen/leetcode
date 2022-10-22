package leetcode

//加油站
//https://leetcode.cn/leetbook/read/top-interview-questions/xmguej/

//这个会超时，有部分测试用例过不去
func canCompleteCircuit(gas []int, cost []int) int {
	//其实核心还是 sum(gas...) >= sum(cost...) 就是满足条件的
	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}
	numsLen := len(gas)
	for i := 0; i < numsLen; i++ {
		//每个为开头去遍历
		sumOfGas := 0
		sumOfCost := 0
		n := 0
		for j := i; n < numsLen; j++ {
			k := j % numsLen
			sumOfGas += gas[k]
			sumOfCost += cost[k]
			if sumOfGas < sumOfCost {
				break
			}
			n++
		}
		if n == numsLen {
			return i
		}
	}
	return -1
}

//其实有的计算是重复的，是可以跳过的
//就像计算之魂里面提到的，其实核心就是想明白哪部分的计算是没有必要的计算
func canCompleteCircuitV2(gas []int, cost []int) int {

	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}
	numsLen := len(gas)
	for i := 0; i < numsLen; {
		//每个为开头去遍历
		sumOfGas := 0
		sumOfCost := 0
		n := 0
		for j := i; n < numsLen; j++ {
			k := j % numsLen
			sumOfGas += gas[k]
			sumOfCost += cost[k]
			if sumOfGas < sumOfCost {
				break
			}
			n++
		}
		if n == numsLen {
			return i
		} else {
			i += n + 1
		}
	}
	return -1
}
