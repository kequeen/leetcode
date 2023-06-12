package leetcode

// 矩阵乘法
func multiplyMatrix(a, b [][]int) [][]int {

	if a == nil || b == nil {
		return nil
	}
	// 首先判断异常,如果a或者b为空，则直接返回nil
	rowA := len(a)
	colA := len(a[0])
	rowB := len(b)
	colB := len(b[0])
	//如果为空，则直接返回nil
	if rowA == 0 || colA == 0 || rowB == 0 || colB == 0 {
		return nil
	}

	//或者是a的列数不等于b的行数，也直接返回nil
	if colA != rowB {
		return nil
	}
	//结果初始化
	res := make([][]int, rowA)
	for i, _ := range res {
		res[i] = make([]int, colB)
	}
	//进行相乘
	//首先考虑a矩阵的遍历
	for i := 0; i < rowA; i++ {
		for j := 0; j < colB; j++ {
			//然后考虑b矩阵的遍历
			for k := 0; k < colA; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}
