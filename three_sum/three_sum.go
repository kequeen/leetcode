package leetcode

import (
	"sort"
)

//https://leetcode.cn/problems/3sum/
//三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

// 你返回所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

//其实两数之和，三数之和，四数之和的解法都还是有差异的，暴力法其实都可以解决，但可以用一些适当的优化降低整体的时间复杂度

//最终还是用到了双指针的方法
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	//最暴力的方法肯定是三层遍历，但三层遍历不能解决重复的问题
	//因为不能重复，所以要通过排序，去掉重复的值
	sort.Ints(nums)
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		//这里也需要判重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//最右边的值
		k := numsLen - 1
		target := -1 * nums[i]
		for j := i + 1; j < numsLen; j++ {
			//如果相同就+1
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			//并且也要判断当前k是否需要向左移动
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j >= k {
				break
			}
			if target == nums[j]+nums[k] {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}

		}
	}
	return ans
}
