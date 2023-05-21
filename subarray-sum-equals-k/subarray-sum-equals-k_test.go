package leetcode

import (
	"testing"
)

func TestSubarrSum(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	if subarraySumV3(nums, k) != 2 {
		t.Fatal("subarraySum error")
	}
	nums2 := []int{1, 2, 3}
	k2 := 3
	if subarraySumV3(nums2, k2) != 2 {
		t.Fatal("subarraySum error")
	}

	num3 := []int{1}
	k3 := 0
	if subarraySumV3(num3, k3) != 0 {
		t.Fatal("subarraySum error")
	}

	num4 := []int{-1, -1, 1}
	k4 := 0
	if subarraySumV3(num4, k4) != 1 {
		t.Fatal("subarraySum error")
	}
}
