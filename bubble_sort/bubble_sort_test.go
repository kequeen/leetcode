package leetcode

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 1, 4}
	bubble_sort(arr)
	fmt.Println(arr)
}
