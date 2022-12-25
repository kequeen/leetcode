package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKList(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 4,
	}
	node3 := &ListNode{
		Val: 5,
	}

	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{
		Val: 1,
	}

	node5 := &ListNode{
		Val: 3,
	}

	node4.Next = node5

	node6 := &ListNode{
		Val: 2,
	}
	node7 := &ListNode{
		Val: 4,
	}
	node6.Next = node7

	lists := []*ListNode{node1, node4, node6}
	//fmt.Println(mergeKLists(lists))
	printList(mergeKLists(lists))

}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}
