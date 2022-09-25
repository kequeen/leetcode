package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTree(t *testing.T) {
	var root1 = &TreeNode{
		Val: 0,
	}
	root1.Left = &TreeNode{
		Val: 5,
	}
	root1.Right = &TreeNode{Val: 6}

	var root2 = &TreeNode{
		Val: 3,
	}
	root2.Left = &TreeNode{
		Val: 5,
	}
	root2.Right = &TreeNode{Val: 6}
	PrintTree(mergeTrees(root1, root2))
	PrintTree(mergeTreesV2(root1, root2))
}

func PrintTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PrintTree(node.Left)
	PrintTree(node.Right)
}
