package leetcode

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	//用最小堆解决
	pq := &PriorityQueue{}
	for _, list := range lists {
		heap.Push(pq, list)
	}
	head := &ListNode{Val: 0}
	tail := head
	for (*pq).Len() != 0 {
		temp := heap.Pop(pq)
		tail.Next = temp.(*item).listnode
		tail = tail.Next
		if temp.(*item).listnode.Next != nil {
			heap.Push(pq, temp.(*item).listnode.Next)
		}
	}
	return head.Next
}

// 最小堆
type item struct {
	val      int
	listnode *ListNode
}

type PriorityQueue []*item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Swap(i, j int) {
	//有时候觉得这种写法也挺奇怪的
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].val < (*pq)[j].val
}

func (pq *PriorityQueue) Push(l interface{}) {
	tempNode := l.(*ListNode)
	//leetcode上需要判断空，是比较奇怪
	if tempNode == nil {
		return
	}
	val := tempNode.Val

	*pq = append((*pq), &item{
		listnode: tempNode,
		val:      val,
	})
}

func (pq *PriorityQueue) Pop() interface{} {
	value := (*pq)[pq.Len()-1]
	(*pq) = (*pq)[:pq.Len()-1]
	return value
}
