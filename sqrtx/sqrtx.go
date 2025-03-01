package sqrtx

// https://leetcode.cn/problems/sqrtx/
// 实现 x 的平方根，返回其整数部分。

func mySqrt(x int) int {
	// 处理特殊情况
	if x == 0 || x == 1 {
		return x
	}

	// 二分查找
	begin := 0
	end := x
	ans := 0
	for begin <= end {
		mid := begin + (end-begin)/2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			end = mid - 1
		} else {
			// 记录可能的整数解
			ans = mid
			begin = mid + 1
		}
	}
	return ans
}
