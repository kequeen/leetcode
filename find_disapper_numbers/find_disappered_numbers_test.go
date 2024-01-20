package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisapperrdNumbers(t *testing.T) {
	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	assert.Equal(t, findDisappearedNumbers(arr), []int{5, 6})
}
