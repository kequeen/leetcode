package leetdoce

//https://leetcode.cn/problems/Ju9Xwi/description/

func leastMinutes(n int) int {
	return calMinutes(1, n)
}

func calMinutes(speed int, leftNum int) int {
	if leftNum <= speed {
		return 1
	}
	return min(leftNum, 1+calMinutes(speed*2, leftNum))
}
