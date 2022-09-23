package leetcode

//计算器
//https://leetcode.cn/leetbook/read/top-interview-questions/xa8q4g/
//其实我理解就是两个栈可以搞定
//还是想复杂了，甚至只需要一个栈，并且这种简单的栈，其实切片完全就可以模拟
func calculate(s string) int {
	result := 0
	stack := []int{}
	//因为数字本身会有多位，需要初始化一个值去处理这种情况
	num := 0
	//默认符号
	preSign := '+'
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		isDigit := false
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			isDigit = true
		}
		//注意最后一位也要处理
		if !isDigit && s[i] != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, num*-1)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			default:
			}
			//存储上一个符号
			preSign = rune(s[i])
			num = 0
		}
	}

	for _, value := range stack {
		result += value
	}

	return result
}
