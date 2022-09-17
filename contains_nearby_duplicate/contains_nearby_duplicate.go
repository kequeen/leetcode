package leetcode

//存在重复元素2
//https://leetcode.cn/problems/contains-duplicate-ii/

//这个方法的边界会有问题，导致 nums[i : i+k] 这种会越界,还是把问题搞复杂了，可以直接用滑动窗口
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	length := len(nums)
// 	//正常情况下
// 	for i := 0; i <= length-k-1; i++ {
// 		if judge(nums[i : i+k+1]) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func judge(slice []int) bool {
// 	dmap := make(map[int]int)
// 	for _, value := range slice {
// 		if dmap[value] == 1 {
// 			return true
// 		} else {
// 			dmap[value] = 1
// 		}
// 	}
// 	return false
// }

//用一下滑动窗口的办法
func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	duplicateMap := make(map[int]bool)
	for i := 0; i < length; i++ {
		if i > k {
			delete(duplicateMap, nums[i-k-1])
		}
		if duplicateMap[nums[i]] {
			return true
		} else {
			duplicateMap[nums[i]] = true
		}

	}
	return false
}
