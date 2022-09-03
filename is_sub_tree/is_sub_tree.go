package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//判断是否子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func dfs(A *TreeNode, B *TreeNode) bool {
	//终止态
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
}
