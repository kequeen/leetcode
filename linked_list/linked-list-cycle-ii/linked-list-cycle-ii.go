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

// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// https://leetcode.cn/problems/linked-list-cycle-ii/description/?favorite=2cktkvj
// 其实就是最简单的快慢指针
// 我反而想了半天，才突然间明白跟之间的链表成环的题目的差别，这个是问的成环的入口
// 之前的题目中，只需要判断有交集，不需要判断具体的成环逻辑。因为交集的地方，其实不一定是成环的地方
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for {
		//如果快指针的下游为nil，则证明没有环
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	//假设 a为到达环前的长度，b为环的长度
	//已经找到相遇点的话，需要再找下成环的点,这个时候 fast = 2s, slow = s
	//fast比slow多走n个环 fast = s + nb => s = nb
	//slow = nb
	//假设走到环的距离 k = a + nb ,目前slow已经走了nb步
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
