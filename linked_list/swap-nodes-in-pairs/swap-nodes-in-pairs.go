package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//两两交换链表中的节点，https://leetcode.cn/problems/swap-nodes-in-pairs/description/
// 但这个目前leetcode那边会导致栈溢出，这个也是递归解法的通病，虽然我一直认为这个优化应当由编译期本身去做这个事情
func swapPairs(head *ListNode) *ListNode {
	//可以拆解为更小的问题去解决
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 采用迭代的方式，而并非递归的方式
func swapPairsV2(head *ListNode) *ListNode {
	//增加头结点，抹平各种异常情况
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
