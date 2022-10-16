package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/
//判断是否平衡二叉树

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)
	if leftDepth-rightDepth < -1 || leftDepth-rightDepth > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

//计算树的高度
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
