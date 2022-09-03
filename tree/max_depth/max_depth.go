package leetcode

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

//求二叉树的最大深度
func maxDepth(root *TreeNode) int {
	//终止条件
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left) + 1
	rightDepth := maxDepth(root.Right) + 1
	return maxVal(leftDepth, rightDepth)
}

func maxVal(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
