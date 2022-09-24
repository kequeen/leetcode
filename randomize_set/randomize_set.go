package leetcode

import "math/rand"

//https://leetcode.cn/leetbook/read/top-interview-questions/xagm3s/
// 实现RandomizedSet 类：

// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

//一般像这种时间复杂度要求O(1)的，其实基本上就只能用map了
//但其实没有想到最后一个条件，其实用纯map的话，是无法做到每个元素有相同的概率被返回的
//所以其实就不可能只用一个map的数据结构，还需要一个切片数组，来支持这个的动态扩展

//还有一点看着解法比较奇妙的是，最终删除元素是使用替换的方式，而不是我自己想的拼接数组的方式，不过确实更为合理

type RandomizedSet struct {
	set map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
		arr: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.arr = append(this.arr, val)
	//存储下标
	this.set[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}
	//首先要移除数组中的这个值，最合适的方式其实是拿另外一个值去覆盖，这里使用数组最后一个值进行覆盖
	last := len(this.arr) - 1
	this.arr[index] = this.arr[last]
	//变更下标
	this.set[this.arr[index]] = index
	this.arr = this.arr[:last]
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.arr))
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
