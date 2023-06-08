package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	assert.Equal(t, 0, root.Val)
	assert.Equal(t, -10, root.Left.Val)
	assert.Equal(t, 5, root.Right.Val)
}
