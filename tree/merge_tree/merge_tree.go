package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//合并二叉树
//https://leetcode.cn/problems/merge-two-binary-trees/?favorite=2cktkvj

//感觉关于二叉树的，基本都需要递归，毕竟不可能去遍历每一个子节点
//这个方法写的太丑了，其实不够优雅，虽然是递归，但很多条件本身没有抹除掉
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	newNode := &TreeNode{}
	if root1 != nil && root2 == nil {
		newNode.Val = root1.Val
		newNode.Left = mergeTrees(root1.Left, nil)
		newNode.Right = mergeTrees(root1.Right, nil)
	} else if root1 == nil && root2 != nil {
		newNode.Val = root2.Val
		newNode.Left = mergeTrees(nil, root2.Left)
		newNode.Right = mergeTrees(nil, root2.Right)
	} else {
		newNode.Val = root1.Val + root2.Val
		newNode.Left = mergeTrees(root1.Left, root2.Left)
		newNode.Right = mergeTrees(root1.Right, root2.Right)
	}
	return newNode
}

//其实整个方法可以优雅很多
//本来合适的算法就应该如此优雅
func mergeTreesV2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	mergeTreesV2(root1.Left, root2.Left)
	mergeTreesV2(root1.Right, root2.Right)
	return root1
}
