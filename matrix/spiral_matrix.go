package leetcode

import "fmt"

//https://leetcode.cn/problems/spiral-matrix/
//螺旋矩阵
//其实就是按照 右 下 左 上 的顺序来进行遍历
//【补充：这个解法越来越复杂，就像在有问题的程序上一直修修补补一般，还是看下标准的方案是怎么解的】
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	rowLen := len(matrix)
	columnLen := len(matrix[0])
	right := columnLen - 1
	down := rowLen - 1
	left := 0
	up := 1
	num := rowLen * columnLen
	column := 0
	row := 0
	for k := 0; k <= num; k++ {
		fmt.Println(k, right, row, column)
		//向右遍历
		for ; column <= right && k <= num; column++ {
			res = append(res, matrix[row][column])
			k++
		}
		column--
		right--
		//向下
		for row = row + 1; row <= down && k <= num; row++ {
			res = append(res, matrix[row][column])
			k++
		}
		row--
		down--
		//向左
		for column = column - 1; column >= left && k <= num; column-- {
			res = append(res, matrix[row][column])
			k++
		}
		column++
		left++
		//向上
		for row = row - 1; row >= up && k <= num; row-- {
			res = append(res, matrix[row][column])
			k++
		}
		row++
		up++
		column++
	}
	return res
}

//打印矩阵
//我觉得按照层去打印会更容易理解一些
func spiralOrderV2(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		row, column              = len(matrix), len(matrix[0])
		order                    = make([]int, row*column)
		left, right, top, bottom = 0, column - 1, 0, row - 1
		index                    = 0
	)
	for left <= right && top <= bottom {
		//从左到右
		for column = left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		//从上到下
		for row = top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		//判断是否还需要遍历
		if left < right && top < bottom {
			//向左遍历
			for column = right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			//向上遍历
			for row = row - 1; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		//最终调整上下左右
		left++
		right--
		top++
		bottom--
	}
	return order
}
