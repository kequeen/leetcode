package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions/xakn6r/
// 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：

// 0 <= i, j, k, l < n
// nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

//这种其实最简单的就能想到的就是遍历，不过那个时间复杂度是要爆炸的 O(N^40
//所以四个数组分别遍历肯定是不符合预期的
//所以思考空间换时间的方式，避免重复计算,这种其实最容易想到的就是map，在O(1)时间内找到自己想要找的值
//这个方法的空间复杂度是O(N),时间复杂度是O(N^2)
//好吧，看了题解，自己空间复杂度想错了，其实是O(N^2),因为可能最坏情况下，v1+v2的值都不相同
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	map1 := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			//存储下这个值有几种可能
			map1[v1+v2]++
		}
	}
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			count += map1[-(v3 + v4)]
		}
	}
	return count
}
