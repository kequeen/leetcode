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

//获取链表的相交
//https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/

//这个解法确实比较巧妙
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	pa := headA
	pb := headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}
