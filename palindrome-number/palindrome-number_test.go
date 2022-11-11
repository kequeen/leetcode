package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	x1 := 121
	assert.Equal(t, isPalindrome(x1), true)

	x2 := -121
	assert.Equal(t, isPalindrome(x2), false)

	x3 := 10
	assert.Equal(t, isPalindrome(x3), false)

}
