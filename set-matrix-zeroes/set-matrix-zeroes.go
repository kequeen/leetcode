package leetcode

// https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 矩阵置零
func setZeroes(matrix [][]int) {
	lenM := len(matrix)
	lenN := len(matrix[0])
	arrM := make([]int, lenM)
	arrN := make([]int, lenN)
	for i := 0; i < lenM; i++ {
		for j := 0; j < lenN; j++ {
			if matrix[i][j] == 0 {
				arrM[i] = 1
				arrN[j] = 1
			}
		}
	}
	//遍历完成之后，按照arrM和arrN进行遍历
	for i := 0; i < lenM; i++ {
		for j := 0; j < lenN; j++ {
			if arrM[i] == 1 || arrN[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
