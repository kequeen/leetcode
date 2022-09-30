package leetcode

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}

	fmt.Println(preorderTraversal(root))
}
