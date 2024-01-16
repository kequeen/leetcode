package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	assert.Equal(t, canPlaceFlowersV2(flowerbed, n), true)
	assert.Equal(t, canPlaceFlowers(flowerbed, n), true)
}
