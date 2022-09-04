package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/problems/add-two-numbers/?favorite=2cktkvj
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 其实我理解这套题目就是大数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//其实可以将链表反转
	reverseL1 := reverseList(l1)
	reverseL2 := reverseList(l2)
	var res *ListNode
	var l3 *ListNode
	carry := 0
	for reverseL1 != nil || reverseL2 != nil {
		current := 0
		if reverseL1 != nil && reverseL2 == nil {
			current = reverseL1.Val + carry
			reverseL1 = reverseL1.Next
		} else if reverseL1 == nil && reverseL2 != nil {
			current = reverseL2.Val + carry
			reverseL2 = reverseL2.Next
		} else {
			current = reverseL1.Val + reverseL2.Val + carry
			reverseL1 = reverseL1.Next
			reverseL2 = reverseL2.Next
		}
		num := current % 10
		carry = current / 10
		node := &ListNode{
			Val: num,
		}
		if res == nil {
			res = node
			l3 = res
		} else {
			l3.Next = node
			l3 = l3.Next
		}
	}

	//新链表算完之后再反转
	return reverseList(res)
}

//如果我再返回一个值，那内存如何分配，这样子考虑下来是否还是原地反转链表更合适
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

func printList(l *ListNode) {
	r := l
	for r != nil {
		fmt.Print(r.Val, " ")
		r = r.Next
	}
}
