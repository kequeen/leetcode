package leetcode

import (
	"fmt"
	"testing"
)

func TestSprialMatrix(t *testing.T) {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrderV2(arr))

	arr2 := [][]int{{2, 3}}
	fmt.Println(spiralOrderV2(arr2))
}
