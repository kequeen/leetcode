package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbersTest(t *testing.T) {
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

	l3 := addTwoNumbers(node1, node4)
	fmt.Println(l3)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
	printList(l3)

}
