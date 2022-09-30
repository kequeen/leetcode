package leetcode

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}

	flatten(root)
	printNode(root)
	//output
	// 0 5 6

}

//只打印右边节点
func printNode(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	printNode(node.Right)
}
