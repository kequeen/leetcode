package leetcode

import (
	"fmt"
	"testing"
)

func TestGetInterSectionNode(t *testing.T) {
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
		Val: 4,
	}

	node4.Next = node2

	interSectionNode := getIntersectionNode(node1, node4)
	fmt.Println(interSectionNode.Val)
}
