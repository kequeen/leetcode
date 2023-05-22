package leetcode

// https://leetcode.cn/problems/two-sum/
// 最经典的两数之和
func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	end := false
	//最简单的遍历
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if end {
				break
			}
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				end = true
				break
			}
		}
	}
	return result
}

// 其实是可以用map的，自己把问题想复杂了
// 以前为啥连map都没想到呢，这个是最简单的解法
func twoSumV2(nums []int, target int) []int {
	targetMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		targetNum := target - nums[i]
		if _, ok := targetMap[targetNum]; ok {
			return []int{targetMap[targetNum], i}
		}
		targetMap[nums[i]] = i
	}
	return nil
}

// 是不是用双指针也可以
// 事实证明我傻了，这又不是有序数组
// 只有有序数据才适合用这种方式
// func twoSumV3(nums []int, target int) []int {
// 	left := 0
// 	right := len(nums) - 1
// 	for left < right {
// 		targetNum := nums[left] + nums[right]
// 		if targetNum == target {
// 			return []int{left, right}
// 		} else if targetNum < target {
// 			left++
// 		} else {
// 			right--
// 		}
// 	}
// 	return nil
// }
