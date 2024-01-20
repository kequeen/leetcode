package leetcode

// 判断是否存在重复元素
// https://leetcode.cn/problems/contains-duplicate/?plan=data-structures&plan_progress=zf360x2
func containsDuplicate(nums []int) bool {
	//用一个map可以解决
	numsMap := make(map[int]bool)
	for _, value := range nums {
		if numsMap[value] {
			return true
		} else {
			numsMap[value] = true
		}
	}
	return false
}
