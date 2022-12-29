package leetcode

//https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/?favorite=2cktkvj
//从前序和中序遍历构造二叉树，还是有点意思的

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

// 所有树的问题，最终都可以归结为递归相关的问题
// 前序遍历 中 左 右
// 中序遍历 左 中 右
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	//根节点构造树
	root := &TreeNode{Val: preorder[0]}
	//寻找出分界线
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	//构造左右节点，问题可以拆解为子问题，寻找左边的前序遍历 + 中序遍历
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
