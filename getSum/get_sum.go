package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions/xmqsae/
//二整数之和,不能使用 + - 等符号
//其实很明显就是考察二进制
func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b << 1
		a = a ^ b
		b = carry
	}
	return a
}
