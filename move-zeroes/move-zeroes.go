package leetcode

//https://leetcode.cn/problems/move-zeroes/description/

// moveZeroes 移动非零元素到数组左侧，最终左边首位元素为最小值
func moveZeroes(nums []int) {
	//双指针的典型题目,不过要注意保持顺序
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			//交换
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
