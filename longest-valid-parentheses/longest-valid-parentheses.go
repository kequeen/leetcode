package leetcode

//https://leetcode.cn/problems/longest-valid-parentheses/description/
//最长有效括号

//开始意识到时动态规划，但动态规划确实又没有想清楚

//最终还是决定用栈的方式去写
//其实就是记录下上次有效的括号的位置，当前如果也有效的话，当前的位置-上次的位置 = 当次的长度
func longestValidParentheses(s string) int {
	maxLen := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				//如果栈已经为空，则说明没有符合条件的，只能继续往后遍历，并且将当前的 ) 的位置给加入进来
				stack = append(stack, i)
			} else {
				//栈不为空，则看下当前的位置-上次的位置，获取长度
				maxLen = max(maxLen, i-stack[len(stack)-1])
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
