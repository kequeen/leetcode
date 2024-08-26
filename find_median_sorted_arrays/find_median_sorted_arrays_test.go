package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMediaSortedArrays(t *testing.T) {
	//验证奇数
	num1 := []int{2, 3, 4, 6}
	num2 := []int{1, 3, 4}
	assert.Equal(t, findMedianSortedArrays(num1, num2), 3.0)

	//验证偶数
	num3 := []int{2, 4, 5}
	assert.Equal(t, findMedianSortedArrays(num2, num3), 3.5)

}
