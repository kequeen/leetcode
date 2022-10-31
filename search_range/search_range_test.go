package leetcode

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))

	nums2 := []int{5, 7, 7, 8, 8, 10}
	target2 := 6
	fmt.Println(searchRange(nums2, target2))
}
