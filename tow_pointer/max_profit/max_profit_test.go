package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	arr := []int{2, 3, 4, 5, 1}
	assert.Equal(t, maxProfit(arr), 3)
	assert.Equal(t, maxProfitV2(arr), 3)
}
