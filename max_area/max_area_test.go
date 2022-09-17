package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))

	fmt.Println(maxAreaV2(nums))

}
