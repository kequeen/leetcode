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

// https://leetcode.cn/problems/kth-node-from-end-of-list-lcci/
// 返回倒数第K个节点
// 下意识就是使用双指针
func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	//遍历结束之后，两个指针一起往前走
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
