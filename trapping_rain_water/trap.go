package leetcode

//https://leetcode.cn/problems/trapping-rain-water/
//终于还是做了熟悉的接雨水
//其实最容易想到的就是，每个节点能接的雨水，是min(往左数的第一个大于它的值, 往右数的第一个大于它的值）- 当前的高度
func trap(height []int) int {
	numsLen := len(height)
	if numsLen == 0 {
		return 0
	}
	sum := 0
	left := make([]int, numsLen)
	right := make([]int, numsLen)
	left[0] = height[0]
	for i := 1; i < numsLen; i++ {
		left[i] = max(left[i-1], height[i])
	}
	right[numsLen-1] = height[numsLen-1]
	for j := numsLen - 2; j >= 0; j-- {
		right[j] = max(right[j+1], height[j])
	}
	for i := 0; i < numsLen; i++ {
		sum = sum + min(left[i], right[i]) - height[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
