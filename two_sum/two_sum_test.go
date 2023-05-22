package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	num := []int{2, 3, 4, 5}
	target := 6
	result := twoSum(num, target)
	assert.Equal(t, result, []int{0, 2})
	assert.Equal(t, twoSumV2(num, target), []int{0, 2})
}
