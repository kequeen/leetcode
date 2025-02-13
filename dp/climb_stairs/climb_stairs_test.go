package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	n := 5
	assert.Equal(t, 8, climbStairs(n))
	assert.Equal(t, 8, climbStairsV2(n))
	assert.Equal(t, 8, climbStairsV3(n))
}
