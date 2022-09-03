package leetcode

//merge合并两个有序数组
//偶尔刷点简单题放松一下
//https://leetcode.cn/leetbook/read/top-interview-questions-easy/xnumcr/
func merge(nums1 []int, m int, nums2 []int, n int) []int {

	result := make([]int, m+n)
	index := 0
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			result[index] = nums2[j]
			j++
		} else {
			result[index] = nums1[i]
			i++
		}
		index++
	}
	if i >= m {
		for j < n {
			result[index] = nums2[j]
			index++
			j++
		}
	}
	if j >= n {
		for i < m {
			result[index] = nums1[i]
			index++
			i++
		}
	}
	return result
}
