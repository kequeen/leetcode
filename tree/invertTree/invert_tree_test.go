package leetcode

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}

	PrintTree(root)

	invertTree(root)

	//output 0 6 5 符合预期
	PrintTree(root)

}

func PrintTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PrintTree(node.Left)
	PrintTree(node.Right)
}
