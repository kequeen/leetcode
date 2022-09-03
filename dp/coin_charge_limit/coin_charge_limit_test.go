package leetcode

import (
	"fmt"
	"testing"
)

func TestCoinChargeLimit(t *testing.T) {
	coins := []int{1, 2, 4, 5}
	money := 13
	fmt.Println(chargeLimit(money, coins))

}
