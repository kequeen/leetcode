package leetcode

import (
	"fmt"
	"testing"
)

func TestChargeCoins(t *testing.T) {
	money := 14
	coins := []int{1, 2, 5}
	fmt.Println(charge(money, coins))
}
