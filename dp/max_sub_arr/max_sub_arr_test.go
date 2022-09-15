package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSubArr(t *testing.T) {
	arr := []int{1, 2, 3, 5, 10, -1, -3}
	fmt.Println(maxSubArray(arr))

}
