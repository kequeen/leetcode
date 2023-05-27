package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	//测试杨辉三角
	result := generate(5)
	assert.Equal(t, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}, result)
	assert.Equal(t, [][]int{{1}}, generate(1))

	assert.Equal(t, [][]int{{1}, {1, 1}}, generateV2(2))
}
