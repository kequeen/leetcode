package leetcode

import "testing"

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	if matrix[1][1] != 0 {
		t.Fatal("setZeroes error")
	}
	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix2)
	if matrix2[0][0] != 0 || matrix2[0][3] != 0 || matrix2[1][0] != 0 || matrix2[2][0] != 0 {
		t.Fatal("setZeroes error")
	}
}
