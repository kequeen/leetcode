package leetcode

import (
	"fmt"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	newHead := swapPairs(head)
	fmt.Println(newHead.Val, newHead.Next.Val)
	if newHead.Val != 2 || newHead.Next.Val != 1 {
		t.Fatal("swapPairs error")
	}
}
