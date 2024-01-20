package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	arr := []int{2, 1, 3, 4, 5, 1, 2}
	assert.Equal(t, false, containsNearbyDuplicate(arr, 2))

	arr2 := []int{9, 9}
	assert.Equal(t, true, containsNearbyDuplicate(arr2, 2))
}
