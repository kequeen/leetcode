package leetcode

import (
	"fmt"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}

	fmt.Println(diameterOfBinaryTree(root))

	fmt.Println(diameterOfBinaryTreeV2(root))
}
