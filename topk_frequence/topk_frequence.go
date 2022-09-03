//https://leetcode.cn/leetbook/read/top-interview-questions/xau4ci/
package leetcode

import (
	"sort"
)

// func topKFrequent(nums []int, k int) []int {
// 	frequentMap := make(map[int]int)
// 	reverseMap := make(map[int]int)
// 	for _, value := range nums {
// 		frequentMap[value]++
// 	}
// 	//互换的方式有问题，value 相同的话会导致被覆盖
// 	frequents := make([]int, 0)
// 	for key, item := range frequentMap {
// 		reverseMap[item] = key
// 		frequents = append(frequents, item)
// 	}
// 	sort.Ints(frequents)
// 	fmt.Println(frequents)

// 	result := make([]int, 0)
// 	for i := len(frequents) - k; i < len(frequents); i++ {
// 		result = append(result, reverseMap[frequents[i]])
// 	}
// 	return result
// }

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	frequentMap := make(map[int]int)
	for _, value := range nums {
		frequentMap[value]++
	}
	//互换的方式有问题，value 相同的话会导致被覆盖
	frequents := make([]int, 0)
	for _, item := range frequentMap {
		frequents = append(frequents, item)
	}
	sort.Ints(frequents)

	result := make([]int, 0)
	//数字不能小于这个数
	minFrequnce := frequents[len(frequents)-k]
	for key, item := range frequentMap {
		if item >= minFrequnce {
			result = append(result, key)
		}
	}
	result = result[0:k]
	return result
}
