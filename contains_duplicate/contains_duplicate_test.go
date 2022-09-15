package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	//判断是否存在重复元素
	nums := []int{1, 2, 4}
	assert.Equal(t, containsDuplicate(nums), false)
}
