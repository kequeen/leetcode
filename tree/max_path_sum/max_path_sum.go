package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum = math.MinInt32

// https://leetcode.cn/leetbook/read/top-interview-questions/x2hnpi/
//动态规划 + 递归
// 不太明白这种写法为什么会一直无法AC
//原来是因为这个原因，不能使用全局变量 https://support.leetcode.cn/hc/kb/article/1194344/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs(root)
	return maxSum
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := max(dfs(node.Left), 0)
	right := max(dfs(node.Right), 0)
	value := left + right + node.Val
	maxSum = max(maxSum, value)
	return max(left, right) + node.Val
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//原来确实是全局变量引入的问题
func maxPathSumV2(root *TreeNode) int {
	var dfs func(*TreeNode) int
	maxSum := math.MinInt32
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		value := left + right + node.Val
		maxSum = max(maxSum, value)
		return max(left, right) + node.Val
	}
	dfs(root)
	return maxSum
}
