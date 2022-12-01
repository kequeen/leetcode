package leetcode

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
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

	printList(node1)
	deleteNode(node2)
	fmt.Println()
	printList(node1)

	//output 符合预期
	//1 2 3
	//1 3

	deleteNodeV2(node1, 3)
	fmt.Println()
	printList(node1)

}

func printList(l *ListNode) {
	r := l
	for r != nil {
		fmt.Print(r.Val, " ")
		r = r.Next
	}
}
