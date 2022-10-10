package leetcode

import "math/rand"

//https://leetcode.cn/leetbook/read/top-interview-questions/xmchc3/
// 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。

// 实现 Solution class:

// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果

//简单分析
//因为要支持 reset，所以肯定需要有个数组存储原本的数据
//然后另外一个要纯随机的话，只能利用 rand 函数，每次 rand 一个0-n之间的数值 rand.Intn函数

//其实这个有点像洗牌算法

type Solution struct {
	arr        []int //原始数组
	shuffleArr []int //洗牌后的数组
}

func Constructor(nums []int) Solution {
	return Solution{
		arr:        append([]int(nil), nums...),
		shuffleArr: nums,
	}
}

//重置之后也是返回洗牌后的数组
func (this *Solution) Reset() []int {
	copy(this.shuffleArr, this.arr)
	return this.shuffleArr
}

func (this *Solution) Shuffle() []int {
	temp := make([]int, len(this.shuffleArr))
	for k, _ := range temp {
		i := rand.Intn(len(this.shuffleArr))
		temp[k] = this.shuffleArr[i]
		this.shuffleArr = append(this.shuffleArr[:i], this.shuffleArr[i+1:]...)
	}
	this.shuffleArr = temp

	return this.shuffleArr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
