package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions/xa6dkt/
//Excel表列序号 其实本质上是一个二十六进制转换的问题
//时间复杂度只遍历一次，是O(N),空间复杂度是O(1)
func titleToNumber(columnTitle string) int {
	num := 0
	for _, value := range columnTitle {
		temp := (value - 'A') + 1
		num = num*26 + int(temp)
	}
	return num
}
