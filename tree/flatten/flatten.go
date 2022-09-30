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

//https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/?favorite=2cktkvj

//二叉树展开为链表
//其实最简单的一个办法就是一个深度优先遍历
// 一般而言，自己都觉得写着复杂的写法，一般不是最佳写法，还有优化空间

func flatten(root *TreeNode) {
	//首先前序遍历完所有节点
	list := preorderTraversal(root)
	//再构造个二叉树,里面的父节点不用动
	//按照这个顺序，重新构建二叉树
	nodeNum := len(list)
	for i := 0; i < nodeNum-1; i++ {
		list[i].Right = list[i+1]
		list[i].Left = nil
	}
}

//前序遍历
func preorderTraversal(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	if root == nil {
		return res
	}
	res = append(res, root)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

//O(1)时间复杂度的还需要再想想，其实应该是每遍历一个左节点，就将它的值插入到右节点上
func flattenV2(root *TreeNode) {
	// cur := root
	// for cur != nil {
	// 	//主要处理左节点不为空的情况
	// 	right = cur.Right
	// 	if cur.Left != nil{
	// 		//将当前的左结点移动到右节点上
	// 		for cur.Right

	// 	}
	// 	//当前的右节点移动到右节点的子节点上
	// }

}
