package leetcode

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	arr := []int{1, 2, 3}
	res := combinationSum(arr, 4)
	for _, value := range res {
		fmt.Println(value)
	}

}
