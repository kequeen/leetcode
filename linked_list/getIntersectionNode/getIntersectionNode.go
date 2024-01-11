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

// 这个解法确实比较巧妙
// 本身就只有两种情况
// 1、不相交，
// A、B互换之后，大家的遍历都是 len(A) + len(B) == len(B) + len(A)，最终遍历完一遍之后也会终止
// 2、相交
// 那大家就会在第二次遍历的时候相遇
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
