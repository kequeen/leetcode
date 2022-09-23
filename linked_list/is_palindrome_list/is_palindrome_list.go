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

//https://leetcode.cn/leetbook/read/top-interview-questions/xaw0rm/
//判断是否回文链表，最简单的想法，其实就是遍历完一遍之后，取一个数组存储所有的数，然后就转换为这个数组是否回文数组的问题

//如果要用O(1)的空间复杂度的话，那就只能反转后半个链表了
func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	//判断数组是否回文
	left := 0
	right := len(arr) - 1

	for left < right {
		if arr[left] == arr[right] {
			left++
			right--
		} else {
			return false
		}
	}

	return true
}
