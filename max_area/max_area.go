package leetcode

//https://leetcode.cn/problems/container-with-most-water/?favorite=2cktkvj
//盛最多水的容器
//暴力法最容易
//但这个方法确实超出了时间限制
func maxArea(height []int) int {
	area := 0
	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			currentArae := min(height[i], height[j]) * (i - j)
			if currentArae > area {
				area = currentArae
			}

		}
	}
	return area
}

//还是要好好想一下双指针的一些场景
//这个的场景其实在于横向肯定是缩小的，而纵向的增大，肯定来源于小的一侧的移动
func maxAreaV2(height []int) int {
	area := 0
	left := 0
	right := len(height) - 1
	for left < right {
		currentArae := min(height[left], height[right]) * (right - left)
		if currentArae > area {
			area = currentArae
		}
		//判断指针向左还是向右移动
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
