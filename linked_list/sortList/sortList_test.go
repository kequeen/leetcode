package leetcode

import (
	"fmt"
	"testing"
)

func TestSortList(t *testing.T) {
	node1 := &ListNode{
		Val: 8,
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
	list2 := sortList(node1)
	printList(list2)

}

//打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
