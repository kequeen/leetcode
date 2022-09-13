package leetcode

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	arr := []int{3, 4, 5, 1, 2, 6, 10, 23}
	if longestConsecutive(arr) != 6 {
		t.Error("最长连续子序列错误，应该为6")
	}
}
