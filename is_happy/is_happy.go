package leetcode

//https://leetcode.cn/problems/happy-number/solutions/
//快乐树
func isHappy(n int) bool {
	//用map存储已经遍历过的数据
	happyMap := make(map[int]bool)
	for n != 1 && !happyMap[n] {
		happyMap[n] = true
		n = calNum(n)

	}

	return n == 1
}

//计算下一个值
func calNum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
