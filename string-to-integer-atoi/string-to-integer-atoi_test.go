package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, 42, myAtoi("    42"))
	assert.Equal(t, -42, myAtoi("  -42"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
}
