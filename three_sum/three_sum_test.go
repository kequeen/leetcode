package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))

	arr2 := []int{0, 0, 0}
	fmt.Println(threeSum(arr2))

	arr3 := []int{1, 2, -2, -1}
	fmt.Println(threeSum(arr3))
}
