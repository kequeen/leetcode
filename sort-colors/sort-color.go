package leetcode

//https://leetcode.cn/problems/sort-colors/?favorite=2cktkvj
//颜色分类
//第一印象这种就是桶排序
//用快排或者冒泡也是一种办法，但这种的时间复杂度最少也为O(NlogN)
//可以用
func sortColors(nums []int) {
	colorMap := make(map[int]int)
	for _, v := range nums {
		colorMap[v]++
	}
	index := 0
	for color := 0; color <= 2; color++ {
		for i := 0; i < colorMap[color]; i++ {
			nums[index] = color
			index++
		}
	}
}

//官方解法各种单指针双指针，反而觉得不如桶排序
