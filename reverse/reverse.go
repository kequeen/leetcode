package leetcode

import "math"

//https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnx13t/
//整数反转，题目不难，但是要注意溢出，主要是还会有个数学论证溢出的问题
func reverse(x int) int {
	res := 0
	for x != 0 {
		//判断下溢出的问题
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		k := x % 10
		res = res*10 + k
		x = x / 10
	}
	return res
}
