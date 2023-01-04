package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	fmt.Println(removeElement(nums2, val2))
}
