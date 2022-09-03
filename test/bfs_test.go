package leetcode

import (
	"fmt"
	"testing"
)

func TestBfs(t *testing.T) {
	var root = initTreeNode()
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}
	fmt.Println(bfs(root))

}
