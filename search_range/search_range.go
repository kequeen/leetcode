package leetcode

//Find First and Last Position of Element in Sorted Array
//https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/?favorite=2cktkvj
func searchRange(nums []int, target int) []int {
	//我的想法是用二分查找找到位置，然后向前和向后遍历
	res := []int{}
	targetLeft := -1
	targetRight := -1
	left := 0
	right := len(nums) - 1
	targetPlace := -1
	for left <= right {
		//开始二分查找，找到对应的位置
		mid := (left + right) / 2
		if nums[mid] == target {
			targetPlace = mid
			break
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if targetPlace != -1 {
		//开始向左和向右遍历
		targetLeft = targetPlace
		targetRight = targetPlace
		for targetLeft >= 0 && nums[targetLeft] == target {
			targetLeft--
		}
		for targetRight < len(nums) && nums[targetRight] == target {
			targetRight++
		}

		//最终结果左右应该都会多算一位
		targetLeft++
		targetRight--

	}
	res = append(res, targetLeft, targetRight)
	return res
}
