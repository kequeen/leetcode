package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t, 3, longestCommonSubsequence("abcde", "ace"))
	assert.Equal(t, 3, longestCommonSubsequenceV2("abcde", "ace"))
}
