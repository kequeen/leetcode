package leetcode

//突然间想明白了那道打印矩阵的题目是如何打印的，当初对 golang 语言本身不熟，反而会耗费一部分精力在语言本身的细节上
//shopee 面试题
/**
Given an integer N, print a matrix in spiral starting with 1 in center.

N = 10

7 8 9 10
6 1 2 11
5 4 3

7 8 9 10
6 1 2 11
5 4 3 12
      13
**/

func printMatrix(num int) interface{} {
	//首先类似于求开方的值，获取结果矩阵的长度，简单一点，从小到大的判断
	matrixLen := 1
	for i := 0; i < num; i++ {
		if i*i >= num {
			matrixLen = i
			break
		}
	}

	matrix := make([][]int, matrixLen)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, matrixLen)
	}
	begin := 0
	//计算起点 如果是偶数，就是/2 - 1 ,奇数就是/2
	if matrixLen%2 == 0 && matrixLen != 1 {
		begin = matrixLen/2 - 1
	} else {
		begin = matrixLen / 2
	}

	right := 1
	down := 1
	left := 2
	up := 2
	row := begin
	column := begin
	matrix[row][column] = 1
	for i := 1; i < num; {
		//向右
		for k := 0; k < right && i < num; k++ {
			column++
			i++
			matrix[row][column] = i
		}
		right++

		//向下
		for k := 0; k < down && i < num; k++ {
			row++
			i++
			matrix[row][column] = i
		}
		down++

		//向左
		for k := 0; k < left && i < num; k++ {
			column--
			i++
			matrix[row][column] = i
		}
		left++

		//向上
		for k := 0; k < up && i < num; k++ {
			row--
			i++
			matrix[row][column] = i
		}

	}

	return matrix

}
