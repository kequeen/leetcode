package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	arr := []int{2, 3, 4, 5, 1}
	fmt.Println(maxProfit(arr))
	fmt.Println(maxProfitV2(arr))
}
