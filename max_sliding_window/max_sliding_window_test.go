package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	arr := []int{2, 3, 4, 1, 3, 2, 7}
	fmt.Println(maxSlidingWindow(arr, 2))
	fmt.Println(maxSlidingWindowV2(arr, 2))
	fmt.Println(maxSlidingWindowV3(arr, 2))

	arr1 := []int{1, 3, 1, 2, 0, 5}
	fmt.Println(maxSlidingWindowV3(arr1, 3))
}
