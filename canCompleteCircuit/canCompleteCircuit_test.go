package leetcode

import (
	"fmt"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
	fmt.Println(canCompleteCircuitV2(gas, cost))
}
