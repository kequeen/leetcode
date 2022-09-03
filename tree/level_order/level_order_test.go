package leetcode

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	result := make([][]int, 0)
	result = append(result, []int{1})
	fmt.Println(result)

	var root = initTreeNode()
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}

	fmt.Println(levelOrder(root))
}
