package leetcode

//https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?favorite=2cktkvj
//Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
//其实就是二分的变种，要想明白二分的条件
//二分查找也不容易啊
func search(nums []int, target int) int {
	//判定二分结束
	if len(nums) == 0 {
		return -1
	}
	//只有一个元素的处理
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}
	left := 0
	right := len(nums) - 1
	//开始标准的二分查找的流程
	for left <= right {
		mid := (left + right) / 2
		//终止
		if target == nums[mid] {
			return mid
		}
		//如果左边是有序的
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				//target在左边这个区间
				right = mid - 1
			} else {
				//target不在这个区间内
				left = mid + 1
			}
		} else {
			//如果右边是有序的
			if target > nums[mid] && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}
