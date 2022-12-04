package leetcode

import (
	"fmt"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowersV2(flowerbed, n))
}
