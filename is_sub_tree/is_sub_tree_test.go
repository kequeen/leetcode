package leetcode

import (
	"fmt"
	"testing"
)

func TestIsSubTree(t *testing.T) {
	node1 := TreeNode{
		Val: 3,
	}
	node2 := TreeNode{
		Val: 4,
	}
	node1.Left = &node2
	node1.Right = nil
	node1.Left.Left = nil
	fmt.Println(isSubtree(&node1, &node2))

	node3 := TreeNode{
		Val: 5,
	}

	fmt.Println(isSubtree(&node1, &node3))
}
