//https://leetcode.cn/leetbook/read/top-interview-questions/xaicbc/
package leetcode

import (
	"container/heap"
)

//kthSmallest 有序矩阵中第K小的元素
//多路归并，只是求这个值的话，用最小堆比较合适
//哪一路的值当前最小，就把哪一路的值继续往后
func kthSmallest(matrix [][]int, k int) int {
	//将第一列的数据压入其中
	minHeap := &Iheap{}
	rowLen := len(matrix)
	columnLen := len(matrix[0])
	for i := 0; i < rowLen; i++ {
		n := node{
			value:  matrix[i][0],
			row:    i,
			column: 0,
		}
		heap.Push(minHeap, n)
	}
	for i := 0; i < k-1; i++ {
		//每次弹出最小值
		minNode := heap.Pop(minHeap).(node)
		if minNode.column != columnLen-1 {
			heap.Push(minHeap, node{matrix[minNode.row][minNode.column+1], minNode.row, minNode.column + 1})
		}
	}

	return heap.Pop(minHeap).(node).value
}

type node struct {
	value  int
	row    int
	column int
}

type Iheap []node

//实现最小堆的各种方法
func (h *Iheap) Push(value interface{}) {
	*h = append(*h, value.(node))
}

func (h *Iheap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func (h Iheap) Len() int {
	return len(h)
}

func (h Iheap) Less(i, j int) bool {
	return h[i].value < h[j].value
}

func (h Iheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
