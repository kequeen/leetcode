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

// 移除重复节点
// https://leetcode.cn/problems/remove-duplicate-node-lcci/
// 其实下意识就是去用map去判断是否重复节点
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nodeMap := make(map[int]bool)
	//初始化
	nodeMap[head.Val] = true
	prev := head
	//去遍历
	for prev.Next != nil {
		cur := prev.Next
		if nodeMap[cur.Val] {
			prev.Next = cur.Next
		} else {
			nodeMap[cur.Val] = true
			prev = prev.Next
		}
	}
	return head
}
