package leetcode

import (
	"container/list"
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := list.New()
	stack.PushBack(2)
	stack.PushBack(3)
	fmt.Println(stack.Back().Value)
}
