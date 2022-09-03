package leetcode

import "container/list"

//关于二叉树的层序遍历
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
			arr = append(arr, node.Value.(*TreeNode).Val)
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
