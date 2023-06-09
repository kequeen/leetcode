package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	assert.Equal(t, 2, searchInsert([]int{1, 3, 5, 6}, 5))
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
}
