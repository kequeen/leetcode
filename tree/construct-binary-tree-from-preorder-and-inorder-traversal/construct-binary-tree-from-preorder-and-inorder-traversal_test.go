package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	PrintTree(root)
}

func PrintTree(node *TreeNode) {
	if node == nil {
		fmt.Println(nil)
		return
	}
	fmt.Println(node.Val)
	PrintTree(node.Left)
	PrintTree(node.Right)
}
