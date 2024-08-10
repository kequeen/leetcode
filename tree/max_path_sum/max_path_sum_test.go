package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	var root = &TreeNode{
		Val: 0,
	}
	assert.Equal(t, maxPathSum(root), 0)
	root.Left = &TreeNode{
		Val: 5,
	}
	root.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: -1}

	assert.Equal(t, maxPathSum(root), 11)
	assert.Equal(t, maxPathSumV2(root), 11)

}
