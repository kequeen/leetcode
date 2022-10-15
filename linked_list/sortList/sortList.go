package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode.cn/leetbook/read/top-interview-questions/xa262d/
//链表排序

func sortList(head *ListNode) *ListNode {
	//最合适的方式是使用归并排序,其实这道题目还是考察挺全面的，考察了如何找到链表的中间节点，以归并排序
	return sort(head, nil)
}

//标准的归并排序
func sort(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//如果只剩一个节点
	if head.Next == tail {
		head.Next = nil
		return head
	}

	//快慢指针寻找中间点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

//merge链表
//因为传进来的是指针，不能改变原来链表中的数据
func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	resListNode := &ListNode{}
	temp, temp1, temp2 := resListNode, list1, list2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 == nil {
		temp.Next = temp2
	} else if temp2 == nil {
		temp.Next = temp1
	}
	return resListNode.Next
}
