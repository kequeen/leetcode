package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChargeCoins(t *testing.T) {
	money := 14
	coins := []int{1, 2, 5}
	assert.Equal(t, 4, charge(money, coins))
}
