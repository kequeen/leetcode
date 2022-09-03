//https://leetcode.cn/leetbook/read/top-interview-questions/xalff2/
//我理解就是有序数组，每插入一个数，插入排序
//最终发现用这种插入排序的方法会超时
//最终还是得用大小堆的方式重新写一遍

//用大小堆的方法终于 AC 了，不过官方还有更为简洁的方法
//https://leetcode.cn/problems/find-median-from-data-stream/solution/shu-ju-liu-de-zhong-wei-shu-by-leetcode-ktkst/

package leetcode

import (
	"container/heap"
)

type MinHeap struct {
	nums []int
}

type MaxHeap struct {
	nums []int
}

//实现 heap要实现 Pop, Push, Len, Swap, Less 五个方法,其余的细节由 container 中的细节所改变
func (mh *MinHeap) Push(value interface{}) {
	mh.nums = append(mh.nums, value.(int))
}

func (mh *MinHeap) Pop() interface{} {
	value := mh.nums[len(mh.nums)-1]
	mh.nums = mh.nums[0 : len(mh.nums)-1]
	return value
}

func (mh *MinHeap) Len() int {
	return len(mh.nums)
}

func (mh *MinHeap) Swap(i, j int) {
	mh.nums[i], mh.nums[j] = mh.nums[j], mh.nums[i]
}

func (mh *MinHeap) Less(i, j int) bool {
	return mh.nums[i] < mh.nums[j]
}

func (mh *MinHeap) Peak() int {
	return mh.nums[0]
}

func (mh *MaxHeap) Push(value interface{}) {
	mh.nums = append(mh.nums, value.(int))
}

func (mh *MaxHeap) Pop() interface{} {
	value := mh.nums[len(mh.nums)-1]
	mh.nums = mh.nums[0 : len(mh.nums)-1]
	return value
}

func (mh *MaxHeap) Len() int {
	return len(mh.nums)
}

func (mh *MaxHeap) Swap(i, j int) {
	mh.nums[i], mh.nums[j] = mh.nums[j], mh.nums[i]
}

func (mh *MaxHeap) Less(i, j int) bool {
	return mh.nums[i] > mh.nums[j]
}

func (mh *MaxHeap) Peak() int {
	return mh.nums[0]
}

type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func Constructor() MedianFinder {
	mf := MedianFinder{
		maxHeap: &MaxHeap{},
		minHeap: &MinHeap{},
	}
	heap.Init(mf.minHeap)
	heap.Init(mf.maxHeap)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	//用一个最大堆和最小堆解决，大堆存放小的数据，小堆存放大的数据
	heap.Push(this.maxHeap, num)
	heap.Push(this.minHeap, heap.Pop(this.maxHeap))

	for this.maxHeap.Len() < this.minHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() < this.maxHeap.Len() {
		return float64(this.maxHeap.Peak())
	}
	return float64(this.minHeap.Peak()+this.maxHeap.Peak()) / 2
}

// type MedianFinder struct {
// 	arr []int
// }

// func Constructor() MedianFinder {
// 	return MedianFinder{
// 		arr: make([]int, 0),
// 	}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	//插入排序
// 	if len(this.arr) == 0 {
// 		this.arr = append(this.arr, num)
// 		return
// 	}
// 	if num <= this.arr[0] {
// 		this.arr = append([]int{num}, this.arr...)
// 		return
// 	}
// 	if num >= this.arr[len(this.arr)-1] {
// 		this.arr = append(this.arr, num)
// 		return
// 	}
// 	//优化的话，可以中位查找，查找到插入位置
// 	insertPlace := 0
// 	// for i := 0; i < len(this.arr)-1; i++ {
// 	// 	if num >= this.arr[i] && num < this.arr[i+1] {
// 	// 		insertPlace = i + 1
// 	// 		break
// 	// 	}
// 	// }
// 	insertPlace = this.binarySearch(1, len(this.arr), num)

// 	newArr := make([]int, 0)
// 	newArr = append(newArr, this.arr[0:insertPlace]...)
// 	newArr = append(newArr, num)
// 	newArr = append(newArr, this.arr[insertPlace:]...)
// 	this.arr = newArr
// 	return
// }

// //采用二分查找
// func (this *MedianFinder) binarySearch(left int, right int, num int) int {
// 	if left >= right {
// 		return left
// 	}
// 	mid := (left + right) / 2
// 	if this.arr[mid] < num {
// 		left = mid + 1
// 		return this.binarySearch(left, right, num)
// 	} else if this.arr[mid] > num {
// 		right = mid
// 		return this.binarySearch(left, right, num)
// 	}
// 	return mid
// }

// func (this *MedianFinder) FindMedian() float64 {
// 	arrLen := len(this.arr)
// 	if arrLen == 0 {
// 		return 0.0
// 	}
// 	if arrLen%2 == 1 {
// 		return float64(this.arr[arrLen/2])
// 	} else {
// 		return float64(this.arr[arrLen/2]+this.arr[arrLen/2-1]) / 2
// 	}
// }

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
