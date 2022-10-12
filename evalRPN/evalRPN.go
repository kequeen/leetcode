package leetcode

import (
	"strconv"
)

//https://leetcode.cn/leetbook/read/top-interview-questions/xaqlgj/

//逆波兰表达式
//其实我理解就是遇到数字就入栈，如果遇到符号，就将栈顶的两个数进行计算
func evalRPN(tokens []string) int {
	//用数组模拟栈
	arr := []int{}
	var n1, n2 int
	result := 0
	for _, v := range tokens {
		arrLen := len(arr)
		switch v {
		case "+":
			n1 = arr[arrLen-1]
			n2 = arr[arrLen-2]
			result = n2 + n1
			arr = arr[:arrLen-2]
			arr = append(arr, result)
		case "-":
			n1 = arr[arrLen-1]
			n2 = arr[arrLen-2]
			arr = arr[:arrLen-2]
			result = n2 - n1
			arr = append(arr, result)
		case "*":
			n1 = arr[arrLen-1]
			n2 = arr[arrLen-2]
			arr = arr[:arrLen-2]
			result = n2 * n1
			arr = append(arr, result)
		case "/":
			n1 = arr[arrLen-1]
			n2 = arr[arrLen-2]
			arr = arr[:arrLen-2]
			// result = int(math.Floor(float64(n2) / float64(n1)))
			result = n2 / n1
			arr = append(arr, result)
		default:
			value, _ := strconv.Atoi(v)
			arr = append(arr, value)
		}
	}
	return arr[0]
}

//官方的写法会更优雅，但其实都是类似的思路，补充下官方的写法
func evalRPNV2(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		value, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, value)
		} else {
			//是符号
			n1 := stack[len(stack)-1]
			n2 := stack[len(stack)-2]
			calValue := 0
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				calValue = n2 + n1
			case "-":
				calValue = n2 - n1
			case "*":
				calValue = n2 * n1
			case "/":
				calValue = n2 / n1
			}
			stack = append(stack, calValue)
		}
	}
	return stack[0]
}
