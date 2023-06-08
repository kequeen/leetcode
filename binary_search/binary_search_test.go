package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	assert.Equal(t, 2, binarySearch(a, 5))
	assert.Equal(t, 1, binarySearch(a, 3))
}
