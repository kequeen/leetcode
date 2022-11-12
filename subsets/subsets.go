package leetcode

//https://leetcode.cn/problems/subsets/?favorite=2cktkvj
//相关链接
func subsets(nums []int) [][]int {
	//其实就是标准的回溯法的应用
	//直接遍历貌似也比较简单
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		size := len(res)
		for j := 0; j < size; j++ {
			temp := append(append([]int{}, res[j]...), nums[i])
			res = append(res, temp)
		}
	}
	return res
}
