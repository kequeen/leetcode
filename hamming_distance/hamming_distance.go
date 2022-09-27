package leetcode

//https://leetcode.cn/problems/hamming-distance/?favorite=2cktkvj
//两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

//这个算法其实还是写复杂了，可以直接用亦或来解决的
//并且没必要用数组来存储，可以直接计算的
func hammingDistance(x int, y int) int {
	//其实就是计算完之后再补全
	arrX := []int{}
	arrY := []int{}
	for x != 0 {
		arrX = append(arrX, x%2)
		x = x / 2
	}
	for y != 0 {
		arrY = append(arrY, y%2)
		y = y / 2
	}
	//对数组补0
	if len(arrX) < len(arrY) {
		for i := len(arrX); i < len(arrY); i++ {
			arrX = append(arrX, 0)
		}
	} else {
		for i := len(arrY); i < len(arrX); i++ {
			arrY = append(arrY, 0)
		}
	}
	diff := 0
	for i := 0; i < len(arrX); i++ {
		if arrX[i] != arrY[i] {
			diff++
		}
	}

	return diff
}

//这个解法就简洁漂亮得多
func hammingDistanceV2(x int, y int) int {
	ans := 0
	for s := x ^ y; s > 0; s = s >> 1 {
		ans += s & 1
	}
	return ans
}
