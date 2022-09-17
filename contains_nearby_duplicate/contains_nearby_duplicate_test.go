package leetcode

import (
	"fmt"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	arr := []int{2, 1, 3, 4, 5, 1, 2}
	fmt.Println(containsNearbyDuplicate(arr, 2))

	arr2 := []int{9, 9}
	fmt.Println(containsNearbyDuplicate(arr2, 2))
}
