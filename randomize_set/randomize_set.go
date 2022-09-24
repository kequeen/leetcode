package leetcode

//https://leetcode.cn/leetbook/read/top-interview-questions/xagm3s/
// 实现RandomizedSet 类：

// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。

//一般像这种时间复杂度要求O(1)的，其实基本上就只能用map了

type RandomizedSet struct {
	set map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.set[val] {
		return false
	}
	this.set[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if this.set[val] {
		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	for key := range this.set {
		return key
	}
	//约定一下，异常返回0
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
