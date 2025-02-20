package leetcode

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}

	fmt.Println(levelOrder(root))
	fmt.Println(levelOrderV2(root))
}
