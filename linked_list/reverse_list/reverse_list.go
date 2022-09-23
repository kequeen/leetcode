package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//最常见的反转链表
//https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
