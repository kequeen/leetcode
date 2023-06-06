package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_quick_sort(t *testing.T) {
	arr := []int{2, 1, 4, 2, 3, 5}
	quick_sort(arr)
	assert.Equal(t, arr, []int{1, 2, 2, 3, 4, 5})
	arr1 := []int{2, 1, 4, 2, 3, 5, 6}
	assert.Equal(t, quickSortV2(arr1), []int{1, 2, 2, 3, 4, 5, 6})

	arr2 := []int{2, 1, 4, 2, 3, 5, 6}
	quickSort(arr2, 0, len(arr2)-1)
	assert.Equal(t, arr2, []int{1, 2, 2, 3, 4, 5, 6})
}
