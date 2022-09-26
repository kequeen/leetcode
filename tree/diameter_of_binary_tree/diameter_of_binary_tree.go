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

//https://leetcode.cn/problems/diameter-of-binary-tree/?favorite=2cktkvj
//二叉树的直径
//感觉是不是要专精二叉树了

//其实应该是递归加简单的动态规划
// diameter[root] = diameter[left] + diameter[right] - 1

//感觉这个写法还是搞得太复杂了
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	//遍历一遍所有节点
	if root == nil {
		return maxDiameter
	}
	//这里这个广度优先遍历还是把问题搞复杂了，其实可以有更简单的写法的，就是遍历
	list := []TreeNode{}
	list = append(list, *root)
	for len(list) != 0 {
		node := list[0]
		if node.Left != nil {
			list = append(list, *node.Left)
		}
		if node.Right != nil {
			list = append(list, *node.Right)
		}
		maxDiameter = max(calDiameter(node.Left)+calDiameter(node.Right), maxDiameter)
		if len(list) > 1 {
			list = list[1:]
		} else {
			list = nil
		}
	}
	return maxDiameter
}

func calDiameter(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(calDiameter(node.Left), calDiameter(node.Right)) + 1
}

//简洁写法
func diameterOfBinaryTreeV2(root *TreeNode) int {
	ans := 1
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		ans = max(l+r+1, ans)
		return max(l, r) + 1
	}
	depth(root)
	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
