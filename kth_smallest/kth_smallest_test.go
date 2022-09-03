package leetcode

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	fmt.Println(kthSmallest(matrix, 8))

}
