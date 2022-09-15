package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 3,
	}

	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{
		Val: 1,
	}

	node5 := &ListNode{
		Val: 8,
	}

	node4.Next = node5

	l := mergeTwoLists(node1, node4)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
