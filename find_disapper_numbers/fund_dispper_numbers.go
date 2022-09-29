package leetcode

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

//https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/?favorite=2cktkvj
func findDisappearedNumbers(nums []int) []int {
	//其实我理解就是一个map的事情
	res := []int{}
	numsMap := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		numsMap[nums[i]] = true
	}
	for i := 1; i <= n; i++ {
		if !numsMap[i] {
			res = append(res, i)
		}
	}
	return res
}
