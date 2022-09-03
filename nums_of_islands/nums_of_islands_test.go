package leetcode

import (
	"testing"
)

func TestNumsOfIslands(t *testing.T) {
	//测试用例1
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if numIslands(grid) != 1 {
		t.Error(numIslands(grid))
	}

	//测试用例2
	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	if numIslands(grid2) != 3 {
		t.Error(numIslands(grid2))
	}

	//测试用例3
	grid3 := [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}
	if numIslands(grid3) != 1 {
		t.Error(numIslands(grid3))
	}

}
