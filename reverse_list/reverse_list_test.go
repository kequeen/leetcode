package leetcode

import (
	"fmt"
	"testing"
)

//关于反转链表的测试
func TestReverseList(t *testing.T) {
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
	//打印数据
	printList(node1)
	fmt.Println()
	rList := reverseList(node1)
	printList(rList)
	//  output
	// 	1 2 3
	//  3 2 1
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
}
