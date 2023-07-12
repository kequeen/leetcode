package leetcode

import "math"

// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

// 函数 myAtoi(string s) 的算法如下：

// 读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
// 返回整数作为最终结果。
// 注意：

// 本题中的空白字符只包括空格字符 ' ' 。
// 除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

func myAtoi(s string) int {
	//其实应该直接调用stronv.Atoi的包的
	//核心其实是找到这个转换的规则
	result := 0
	if len(s) == 0 {
		return 0
	}
	sign := 1
	newS := s
	spaceNum := 0
	overFlowFlag := false
	for i, _ := range newS {
		if isSpace(rune(newS[i])) {
			spaceNum++
		} else {
			break
		}
	}
	newS = newS[spaceNum:]
	if newS[0] == '-' {
		sign = -1
		newS = newS[1:]
	} else if newS[0] == '+' {
		newS = newS[1:]
	}
	for i, _ := range newS {
		if isSpace(rune(newS[i])) {
			continue
		}
		if newS[i] >= '0' && newS[i] <= '9' {
			num := newS[i] - '0'
			result = result*10 + int(num)
			//判断是否溢出
			if overFlow(sign * result) {
				overFlowFlag = true
				break
			}
		} else {
			//有特殊字符就直接终止
			break
		}
	}
	if overFlowFlag {
		if sign == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	return sign * result
}

func isSpace(i rune) bool {
	return i == ' '
}

// 判断一个数是否溢出
// 如果不能用这种方式，还需要用比较复杂的判断
func overFlow(num int) bool {
	//最大值 最小值
	if num < math.MaxInt32 && num > math.MinInt32 {
		return false
	}
	return true
}
