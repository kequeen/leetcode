package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)

}

func helper(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (right + left) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
