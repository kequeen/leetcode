package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	assert.Equal(t, longestValidParentheses(s), 2)

	s1 := ")()())"
	assert.Equal(t, longestValidParentheses(s1), 4)

	s2 := ""
	assert.Equal(t, longestValidParentheses(s2), 0)
}
