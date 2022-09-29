package leetcode

//https://leetcode.cn/problems/counting-bits/?favorite=2cktkvj
//比特币计数
//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
func countBits(n int) []int {
	ans := []int{}
	for i := 0; i <= n; i++ {
		temp := i
		count := 0
		for temp != 0 {
			if temp%2 == 1 {
				count++
			}
			temp = temp / 2
		}
		ans = append(ans, count)
	}
	return ans
}
