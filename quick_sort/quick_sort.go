package leetcode

func quick_sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	//中间数的位置
	mid := arr[0]
	//左右两边的游标
	head, tail := 0, len(arr)-1
	//从左到右遍历
	for i := 1; i <= tail; {
		//如果小于中间数，放在左边
		if arr[i] < mid {
			arr[head], arr[i] = arr[i], arr[head]
			head++
			i++
		} else {
			//如果大于中间数，放在右边
			arr[i], arr[tail] = arr[tail], arr[i]
			tail--
		}
	}
	//递归左边
	quick_sort(arr[:head])
	//递归右边
	quick_sort(arr[head+1:])
}

// 最容易理解的快速排序的写法，返回一个新数组
// 这种方式最容易理解，但唯一的问题就是需要不断的创建新数组
func quickSortV2(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	temp := arr[0]
	var left []int
	var right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < temp {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	//递归再对左右数组进行排序
	left = quickSortV2(left)
	right = quickSortV2(right)
	return append(append(left, temp), right...)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		//寻找切割的下标
		pivotIndex := partition(arr, low, high)
		//对左右再分别排序
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}

}

// 分区函数
func partition(arr []int, low int, high int) int {
	//取一个数作为基准
	privot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < privot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	//交换下位置
	arr[high], arr[i+1] = arr[i+1], arr[high]
	return i + 1
}
