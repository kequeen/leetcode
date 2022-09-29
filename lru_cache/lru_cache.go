package leetcode

//其实我理解还是一个map + array 可以解决
//https://leetcode.cn/problems/lru-cache/
//时间复杂度 O(n) 空间复杂度也是O(n)
type LRUCache struct {
	arr      []int
	cacheMap map[int]int
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		arr:      []int{},           //存 index, key
		cacheMap: make(map[int]int), //直接存key，value
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	_, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	//get的时候，也要移动其到数组的最前
	index := 0
	for i := 0; i < len(this.arr); i++ {
		if key == this.arr[i] {
			index = i
			break
		}
	}
	//这个判断还是太容易搞错了，就应该从最后一位开始遍历
	for i := index; i > 0; i-- {
		this.arr[i] = this.arr[i-1]
	}
	this.arr[0] = key
	return this.cacheMap[key]
}

func (this *LRUCache) Put(key int, value int) {
	//首先判断是否已经在

	_, ok := this.cacheMap[key]
	if !ok {
		//如果不在，那就直接加进去，加在第一位
		this.arr = append([]int{key}, this.arr...)
	} else {
		//如果在,那就直接移动下位置到第一位，然后第一位的全部后移
		//寻找index的位
		index := 0
		for i := 0; i < len(this.arr); i++ {
			if key == this.arr[i] {
				index = i
				break
			}
		}
		for i := index; i > 0; i-- {
			this.arr[i] = this.arr[i-1]
		}
		this.arr[0] = key
	}

	//然后判断数组长度，如果超过capacity,就把最后一位移出
	if len(this.arr) > this.capacity {
		//删除掉最后一位元素的值
		delete(this.cacheMap, this.arr[len(this.arr)-1])
		this.arr = this.arr[:len(this.arr)-1]
	}
	this.cacheMap[key] = value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
