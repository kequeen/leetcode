package leetcode

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

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

	intersectionNode := getIntersectionNode(node1, node4)
	fmt.Println(intersectionNode.Val)

}
