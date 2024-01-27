package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindJudge(t *testing.T) {

	n := 2
	trust := [][]int{{1, 2}}
	assert.Equal(t, findJudge(n, trust), 2)

	n2 := 3
	trust2 := [][]int{{1, 3}, {2, 3}}
	assert.Equal(t, findJudge(n2, trust2), 3)

	n3 := 3
	trust3 := [][]int{{1, 3}, {2, 3}, {3, 1}}
	assert.Equal(t, findJudge(n3, trust3), -1)

}
