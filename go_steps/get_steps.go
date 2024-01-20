package leetcode

// 给定一个正整数 x, 可以对x做如下变换:
// x 为奇数时，x => x-1 or x => x+1
// x 为偶数时，x => x/2
// 则至少对x做多少次变换可以变换为1

var dpMap = make(map[int]int)

func GetSteps(x int) int {
	if x == 0 || x == 1 {
		return 0
	}
	if dpMap[x] != 0 {
		return dpMap[x]
	}
	count := 0
	if x%2 == 0 {
		//偶数
		count = GetSteps(x/2) + 1
	} else {
		count = min(GetSteps(x-1), GetSteps(x+1)) + 1
	}
	dpMap[x] = count
	return dpMap[x]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
