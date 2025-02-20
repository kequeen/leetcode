package leetcode

import "container/list"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   interface{}
}

// 关于二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return nil
	}
	//去进行BFS,但是会比较麻烦一些，需要考虑层高
	p := list.New()
	p.PushBack(root)
	for p.Len() != 0 {
		q := list.New()
		arr := make([]int, 0)
		for p.Len() != 0 {
			node := p.Front()
			if node.Value.(*TreeNode).Left != nil {
				q.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil {
				q.PushBack(node.Value.(*TreeNode).Right)
			}
			arr = append(arr, node.Value.(*TreeNode).Val.(int))
			p.Remove(node)
		}
		if len(arr) != 0 {
			result = append(result, arr)
		}
		//每一次取完当前所有的值，然后把下一层的值放在另一个数组里面
		p = q
	}

	return result
}

// 其实可以不需要list这个结构，就正常用BFS就可以了
func levelOrderV2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result := make([][]int, 0)
	for len(queue) > 0 {
		arr := make([]int, 0)
		for _, node := range queue {
			arr = append(arr, node.Val.(int))
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		result = append(result, arr)
	}
	return result
}
