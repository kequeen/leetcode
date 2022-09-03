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
