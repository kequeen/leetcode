//https://leetcode.cn/leetbook/read/top-interview-questions/xatgye/
package leetcode

//这种方式重复计算比较多，果然又超时了
func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen < k {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	maxNum := nums[0]
	result := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {
		for j := i; j < i+k; j++ {
			if maxNum < nums[j] {
				maxNum = nums[j]
			}
		}
		result = append(result, maxNum)
		maxNum = nums[i+1]
	}
	return result
}

//优化版解法，一般优化其实核心还是在于减少重复计算
func maxSlidingWindowV2(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen < k {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	maxNum := nums[0]
	result := make([]int, 0)
	for i := 0; i < numsLen; i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
		if i < k-1 {
			continue
		}
		result = append(result, maxNum)
		//判断最前面一位是否最大值
		if maxNum == nums[i-k+1] {
			//需要重新选一个最大值
			maxNum = nums[i-k+2]
			for j := i - k + 3; j <= i; j++ {
				if maxNum < nums[j] {
					maxNum = nums[j]
				}
			}
		}
	}
	return result
}

//其实可以用单调栈的方式来进行解决
//看了半天才突然间想明白为什么要存储下标,因为一直变小的话，其实没有办法只根据大小判断来退栈
//所以需要存储下标来兜底
func maxSlidingWindowV3(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen < k {
		return []int{}
	}
	if k == 1 {
		return nums
	}
	result := make([]int, 0)
	//构造单调栈
	stack := []int{}
	push := func(i int) {
		//如果大于最初的值，那么就可以将栈底的值去掉
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}

	for j := 0; j < k; j++ {
		push(j)
	}

	result = append(result, nums[stack[0]])

	for j := k; j < numsLen; j++ {
		push(j)
		//判断下是否已经超了当前窗口
		for j-k >= stack[0] {
			stack = stack[1:]
		}

		result = append(result, nums[stack[0]])
	}
	return result
}
