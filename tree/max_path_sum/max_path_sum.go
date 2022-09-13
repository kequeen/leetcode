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
	return max(left+node.Val, right+node.Val)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
