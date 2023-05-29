package leetcode

//滑动窗口的平均值
// 输入：
// inputs = ["MovingAverage", "next", "next", "next", "next"]
// inputs = [[3], [1], [10], [3], [5]]
// 输出：
// [null, 1.0, 5.5, 4.66667, 6.0]

// 解释：
// MovingAverage movingAverage = new MovingAverage(3);
// movingAverage.next(1); // 返回 1.0 = 1 / 1
// movingAverage.next(10); // 返回 5.5 = (1 + 10) / 2
// movingAverage.next(3); // 返回 4.66667 = (1 + 10 + 3) / 3
// movingAverage.next(5); // 返回 6.0 = (10 + 3 + 5) / 3

// 为什么有时候觉得理解题目更难
// https://leetcode.cn/problems/qIsx9U/
// 往回看，还发现过去写的有的答案是错的
type MovingAverage struct {
	size int
	sum  int
	arr  []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.arr) == this.size {
		//干掉首位的
		this.sum = this.sum - this.arr[0]
		this.arr = this.arr[1:]
	}
	this.arr = append(this.arr, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.arr))
}
