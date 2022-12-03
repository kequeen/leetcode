package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 2,
	}

	node1.Next = node2
	node2.Next = node3

	printList(removeDuplicateNodes(node1))

}

func printList(l *ListNode) {
	r := l
	for r != nil {
		fmt.Print(r.Val, " ")
		r = r.Next
	}
}
