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
		Val: 1,
	}

	node1.Next = node2
	node2.Next = node3

	fmt.Println(isPalindrome(node1))

}
