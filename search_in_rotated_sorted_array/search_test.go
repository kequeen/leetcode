package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	arr1 := []int{4, 5, 6, 7, 0, 1, 2}
	target1 := 0
	assert.Equal(t, search(arr1, target1), 4)

	nums2 := []int{4, 5, 6, 7, 0, 1, 2}
	target2 := 3
	assert.Equal(t, search(nums2, target2), -1)

	nums3 := []int{5, 1, 3}
	target3 := 3
	assert.Equal(t, search(nums3, target3), 2)
}
