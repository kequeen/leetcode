package leetcode

import (
	"fmt"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	arr := []int{2, 0, 1}
	fmt.Println(countSmaller(arr))

	fmt.Println(countSmallerV2(arr))

}
