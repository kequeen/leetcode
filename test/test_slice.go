package leetcode

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	//测试下切片相关的操作
	var a []int
	if a == nil {
		fmt.Println("a is nil")
	}

	b := make([]int, 0)
	if b == nil {
		fmt.Println("b is nil")
	}
	//b = append(b, 1)
	fmt.Println(b)
}
