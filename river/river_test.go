package leetcode

import (
	"fmt"
	"testing"
)

func TestCanCrossRiver(t *testing.T) {
	arr1 := []int{0, 1, 3, 6, 10, 15}
	fmt.Println(canCrossRiver(arr1))

	arr2 := []int{0, 1, 3, 7, 10, 15}
	fmt.Println(canCrossRiver(arr2))

	arr3 := []int{0, 1, 3, 6, 7, 10, 15}
	fmt.Println(canCrossRiver(arr3))
}
