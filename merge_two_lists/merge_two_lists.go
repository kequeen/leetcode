package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表 确实有点归并排序的意思
// https://leetcode.cn/problems/merge-two-sorted-lists/?favorite=2cktkvj
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h := &ListNode{}
	l := h
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			l.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			l.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		l = l.Next
	}

	// 类似于归并排序，直接判断l1 或者 l2已经遍历完，剩下的直接合并完成
	if list1 != nil {
		l.Next = list1
	}

	if list2 != nil {
		l.Next = list2
	}

	return h.Next
}
