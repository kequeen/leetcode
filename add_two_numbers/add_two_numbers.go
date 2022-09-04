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
	var res *ListNode
	var l3 *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		current := 0
		if l1 != nil && l2 == nil {
			current = l1.Val + carry
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			current = l2.Val + carry
			l2 = l2.Next
		} else {
			current = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
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
	if carry != 0 {
		l3.Next = &ListNode{
			Val: carry,
		}
	}
	return res
}

func printList(l *ListNode) {
	r := l
	for r != nil {
		fmt.Print(r.Val, " ")
		r = r.Next
	}
}
