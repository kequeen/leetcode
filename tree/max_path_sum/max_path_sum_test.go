package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	fmt.Println(maxPathSum(root))
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: -1}

	fmt.Println(maxPathSum(root))

	fmt.Println(maxPathSumV2(root))

}
