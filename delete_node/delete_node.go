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
