package leetcode

// https://leetcode.cn/problems/can-place-flowers/
// 种花问题

// 没想到最合适的方法其实是利用贪心去解决
// 根据上一个花盆的位置和下一个花盆的位置，可以计算出中间应该有多少花盆
// 其实是分为两种情况，1、左边没种花，2、左边种了花
// 贪心的一些标准过程，先求出局部最优解，

// 问了一下chatgpt，还有更为简单的解法
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	count := 0

	for i := 0; i < length; i += 2 {
		if flowerbed[i] == 0 {
			// 当前位置是0，检查下一个位置
			if i+1 == length || flowerbed[i+1] == 0 {
				// 可以在当前位置或下一个位置种花
				count++
			} else {
				// 跳过下一个位置，因为相邻的地块不能同时种花
				i++
			}
		}
	}

	return count >= n
}

func canPlaceFlowersV2(flowerbed []int, n int) bool {
	count := 0
	//优化的话，其实就是比较下当前计算的count，是否已经达到预期的目标
	//左边界
	prev := -1
	m := len(flowerbed)
	//先计算已经种花的空隙是否能够插入新的花
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			//计算下是否能放花
			if prev < 0 {
				//左边没种花
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	// 没有种花的地方是否还能放花
	if prev < 0 {
		//如果整个都没有种花
		count += (m + 1) / 2
	} else {
		//注意一下右边界的情况，这个等价于右边没有花
		count += (m - prev - 1) / 2
	}
	return count >= n
}

//这种问题的主要迭代公式搞清楚了，但是还是会有很多的边界的处理需要注意，其实这个应该也是一种面试的考察
