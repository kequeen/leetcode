package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	arr := []int{2, 3, 4, 1, 3, 2, 7}
	fmt.Println(maxSlidingWindow(arr, 2))
}
