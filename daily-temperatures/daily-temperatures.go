package leetcode

// https://leetcode.cn/problems/daily-temperatures/?favorite=2cktkvj
// 每日温度
// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。
func dailyTemperatures(temperatures []int) []int {
	//这种应该下意识会有O(n)的解决方案
	//没想到是用这种单调栈的方式,栈顶的元素永远是最小的，遇到比他大的数的时候，开始推出栈内的元素
	ans := make([]int, len(temperatures))
	stack := []int{}
	for i := 0; i < len(temperatures); i++ {
		//比较里面的值
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
