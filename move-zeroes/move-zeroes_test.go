package leetcode

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	result := []int{1, 3, 12, 0, 0}
	if !compareTwoArray(nums, result) {
		t.Fatal("moveZeroes error")
	}
}

// 比较两个数组相等
func compareTwoArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
