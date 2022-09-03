package leetcode

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BFS 遍历节点
func bfs(root *TreeNode) []int {
	//其实就是一个队列,先进先出
	res := make([]int, 0)

	if root == nil {
		return res
	}
	//首先把父节点放进去
	queue := list.New()
	queue.Init()
	queue.PushBack(root)
	for queue.Len() != 0 {
		temp := queue.Front()
		res = append(res, temp.Value.(*TreeNode).Val)
		//新增节点
		if temp.Value.(*TreeNode).Left != nil {
			queue.PushBack(temp.Value.(*TreeNode).Left)
		}
		if temp.Value.(*TreeNode).Right != nil {
			queue.PushBack(temp.Value.(*TreeNode).Right)
		}
		queue.Remove(temp)
	}
	return res
}
