package leetcode

// https://leetcode.cn/problems/unique-binary-search-trees/description/?favorite=2cktkvj
// 不同的二叉搜索树，刷完感觉像是数学题目
// 其实最核心的还是问题拆解，从1个节点，2个节点，3个节点，一步步分解出来
// 每个节点考虑左边的节点，右边的节点等各种情况
func numTrees(n int) int {
	//ans := sum（ 左边的子解 * 右边的子解）
	ans := make([]int, n+1)
	ans[0] = 1
	ans[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			ans[i] += ans[j-1] * ans[i-j]
		}
	}
	return ans[n]
}
