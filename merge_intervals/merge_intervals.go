package leetcode

import "sort"

//https://leetcode.cn/problems/merge-intervals/?favorite=2cktkvj
//合并区间
//只有特别熟练，慢慢才能不害怕各种题目

//我理解分为两步，
//第1步是将所有的排序，按照左边的数字从小到大进行排序
//第2步是从左到右遍历，有当前的左边界和右边界，假如下一个词的左边界 > 当前的右边界，就生成新的区间
func merge(intervals [][]int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return res
	}
	//sortArr(intervals)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	//我们认为格式是可信的
	left := intervals[0][0]
	right := intervals[0][1]

	for _, arr := range intervals {
		//如果左边大于原来的最右区间，则更新
		if arr[0] > right {
			res = append(res, []int{left, right})
			left = arr[0]
			right = arr[1]
		} else {
			//比较下右边区间的大小,看是否需要更新
			if arr[1] > right {
				right = arr[1]
			}
		}
	}
	//遍历完需要把最后的这个区间补上
	res = append(res, []int{left, right})
	return res
}

//需要对数组进行排序
func sortArr(arr [][]int) {
	//就正常的冒泡排序,从小到大排序
	lenArr := len(arr)
	for i := 1; i < lenArr; i++ {
		for j := 0; j < i; j++ {
			if arr[i][0] < arr[j][0] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
