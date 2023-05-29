package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingAverageTest(t *testing.T) {
	obj := Constructor(3)
	assert.Equal(t, 1.0, obj.Next(1))
	assert.Equal(t, 5.5, obj.Next(10))
	assert.Equal(t, 4.666666666666667, obj.Next(3))
	assert.Equal(t, 6.0, obj.Next(5))
}
