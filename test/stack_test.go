package leetcode

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	// stack := list.New()
	// stack.PushBack(2)
	// stack.PushBack(3)
	// fmt.Println(stack.Back().Value)

	//测试下golang中的hash，看看是否是类似于java中的tostring和hashcode的方式
	a := TreeNode{Val: 2}
	b := TreeNode{Val: 2}
	arr := make(map[any]int)
	arr[a] = 1
	arr[b] = 2
	for k, v := range arr {
		fmt.Println(k, v)
	}
	fmt.Println(len(arr))
}
