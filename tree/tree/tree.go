package leetcode

type treeNode struct {
	left  *treeNode
	right *treeNode
	val   interface{}
}

type TreeNode struct {
	left  *treeNode
	right *treeNode
	val   interface{}
}

func nodeInit() *treeNode {
	return &treeNode{}
}
