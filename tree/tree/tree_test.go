package leetcode

import (
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	root := nodeInit()
	root.left = &treeNode{
		val: 3,
	}
	root.right = &treeNode{
		val: 4,
	}

	fmt.Println(root.left.val)

}
