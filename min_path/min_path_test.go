package leetcode

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(minPathSum(grid))
}
