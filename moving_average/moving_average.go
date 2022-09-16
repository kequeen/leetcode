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

//为什么有时候觉得理解题目更难
//https://leetcode.cn/problems/qIsx9U/
type MovingAverage struct {
	length    int
	maxLength int
	arr       []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		length:    0,
		maxLength: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.length < this.maxLength {
		this.length++
	}
	this.arr = append(this.arr, val)
	arrLen := len(this.arr)
	sum := 0
	for i := arrLen - 1; i > arrLen-1-this.length; i-- {
		sum += this.arr[i]
	}
	return float64(sum) / float64(this.length)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
