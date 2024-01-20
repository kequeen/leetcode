package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSum(t *testing.T) {
	a := 3
	b := 4
	assert.Equal(t, getSum(a, b), 7)
}
