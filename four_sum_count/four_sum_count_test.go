package leetcode

import (
	"fmt"
	"testing"
)

func TestFourSumCount(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
