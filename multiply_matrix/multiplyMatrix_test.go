package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplyMatrix(t *testing.T) {
	assert.Equal(t, [][]int(nil), multiplyMatrix([][]int{{}}, [][]int{{}}))
	assert.Equal(t, [][]int(nil), multiplyMatrix([][]int{{1}}, [][]int{{}}))
	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{3}, {4}}
	assert.Equal(t, [][]int{{11}, {25}}, multiplyMatrix(a, b))
}
