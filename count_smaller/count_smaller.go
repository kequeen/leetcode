package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions/xajl22/
//计算右侧小于当前元素的个数

//这种是暴力方法，时间复杂度是O(N^2)，目前超时了
func countSmaller(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	if numsLen == 1 {
		return result
	}
	//从后往前遍历，可以复用前面计算的结果
	for i := numsLen - 2; i >= 0; i-- {
		for j := i + 1; j < numsLen; j++ {
			if nums[j] < nums[i] {
				result[i] = result[i] + 1
			}
		}
	}
	return result
}

//看到标准解法里面其实是构造了一种特殊的数据结构去承载这种情况
func countSmallerV2(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, numsLen)
	return result
}
