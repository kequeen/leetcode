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

//删除链表中的节点
//https://leetcode.cn/leetbook/read/top-interview-questions/xadve1/
//其实一直比较有疑惑的点是，如何去比较两个节点是否相同
//今晚跑步的时候突然想明白了，其实这个都没有给你头节点，其实解法就肯定不是从头部遍历
//并且之前面试遇到的如何判断一颗树是另外一颗树的子树的问题,其实当节点的值相同时，要遍历递归其每个子节点，这样才能判断是否相同
//上次判断地址的方式，是不是把别人都惊呆了。。

func deleteNode(node *ListNode) {
	if node.Next == nil {
		node = nil
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// 第二种版本的删除链表
// https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// 需要删除链表中的值
func deleteNodeV2(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	prev := head
	for prev != nil {
		if prev.Val == val {
			if prev.Next != nil {
				//删除节点
				prev.Val = prev.Next.Val
				prev.Next = prev.Next.Next
			} else {
				//稍微有点想不明白空值为什么不行
				prev = nil
			}
			break
		}
		prev = prev.Next
	}
	return head
}

// 双指针解法
func deleteNodeV3(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	prev := head
	cur := head.Next
	for cur != nil && cur.Val != val {
		prev = cur
		cur = cur.Next
	}
	//就算是最后一位也不例外
	if cur != nil {
		prev.Next = cur.Next
	}
	return head
}
