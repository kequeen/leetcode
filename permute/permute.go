package leetcode

//https://leetcode.cn/problems/permutations/
//全排列，其实就是回溯，这块一直也没有吃透
//对于一些传值与传引用的地方也是没有完全理解清楚

func permute(nums []int) [][]int {
	var res = make([][]int, 0)
	var backtrace func([]int, int, int)
	backtrace = func(nums []int, i int, length int) {
		if i == length {
			newNum := append([]int{}, nums...)
			res = append(res, newNum)
		}
		for j := i; j < length; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			backtrace(nums, i+1, length)
			nums[i], nums[j] = nums[j], nums[i]
		}

	}
	backtrace(nums, 0, len(nums))
	return res
}

//回溯法
