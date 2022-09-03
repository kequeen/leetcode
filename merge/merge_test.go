package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {

	a := []int{1, 2, 3, 4, 7, 0, 0, 0}
	b := []int{2, 5, 10}
	merge(a, 5, b, 3)
	fmt.Println(a)
	sort.Ints(a)
	c := []int{4, 5, 1, 2, 3, 4, 5, 98}
	fmt.Println(ShellSort(c))

}

func ShellSort(nums []int) []int {
	//外层步长控制
	for step := len(nums) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(nums); i += step {
			//满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
	return nums
}
