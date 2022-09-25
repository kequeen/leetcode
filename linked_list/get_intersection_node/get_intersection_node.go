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

//相交链表
//https://leetcode.cn/problems/intersection-of-two-linked-lists/?favorite=2cktkvj
//其实最简单的方式，就是从左到右遍历一遍，然后反转取其交集，从后往前遍历，不同的就返回上个节点
//不过这个解法确实太蠢了，还是看了答案的方法比较妙
//方法一，利用map来判断这个，其实跟我上次判断子树的很类似，我是觉得合法性值得怀疑
//方法二的双指针确实被惊艳到了

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa := headA
	pb := headB
	//其实用了比较巧妙的证明
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
