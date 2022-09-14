package leetcode

import (
	"fmt"
	"testing"
)

func TestIsLandPermeter(t *testing.T) {
	grid3 := [][]int{{1, 1, 1}, {0, 1, 0}, {1, 1, 1}}

	fmt.Println(islandPerimeter(grid3))
}
