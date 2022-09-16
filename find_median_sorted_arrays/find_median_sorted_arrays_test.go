package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMediaSortedArrays(t *testing.T) {
	//验证奇数
	num1 := []int{2, 3, 4, 6}
	num2 := []int{1, 3, 4}
	fmt.Println(findMedianSortedArrays(num1, num2))

	//验证偶数
	num3 := []int{2, 4, 5}
	fmt.Println(findMedianSortedArrays(num2, num3))

}
