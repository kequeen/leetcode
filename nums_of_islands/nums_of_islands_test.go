package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	// if numIslandsV2(grid) != 1 {
	// 	t.Error(numIslandsV2(grid))
	// }

	assert.Equal(t, numIslandsV3(grid), 1)

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
	// if numIslandsV2(grid2) != 3 {
	// 	t.Error(numIslandsV2(grid2))
	// }

	assert.Equal(t, numIslandsV3(grid2), 3)

	//测试用例3
	grid3 := [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}
	if numIslands(grid3) != 1 {
		t.Error(numIslands(grid3))
	}

	// if numIslandsV2(grid3) != 1 {
	// 	t.Error(numIslandsV2(grid2))
	// }

	assert.Equal(t, numIslandsV3(grid3), 1)

}
