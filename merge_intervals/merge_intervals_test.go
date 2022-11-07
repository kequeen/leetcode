package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	arr1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(arr1))

	arr2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(arr2))
}
