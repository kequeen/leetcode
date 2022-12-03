package leetcode

import (
	"fmt"
	"testing"
)

func TestKthToLast(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
	}
	fmt.Println(kthToLast(node1, 1))

	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 3,
	}

	node1.Next = node2
	node2.Next = node3

	fmt.Println(kthToLast(node1, 2))

}
