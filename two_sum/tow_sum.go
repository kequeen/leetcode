package leetcode

//https://leetcode.cn/problems/two-sum/
//最经典的两数之和
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

//其实是可以用map的，自己把问题想复杂了
