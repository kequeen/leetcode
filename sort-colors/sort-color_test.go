package leetcode

import (
	"fmt"
	"testing"
)

func TestSortColor(t *testing.T) {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}
