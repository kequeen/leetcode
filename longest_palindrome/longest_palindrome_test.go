package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongPalindrome(t *testing.T) {
	s := "abcdd"
	assert.Equal(t, longestPalindrome(s), "dd")
	s = "cbbd"
	assert.Equal(t, longestPalindrome(s), "bb")
}
