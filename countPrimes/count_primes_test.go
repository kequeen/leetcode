package leetcode

import (
	"fmt"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	fmt.Println(countPrimes(4))
	fmt.Println(countPrimesV2(4))
	fmt.Println(countPrimes(5))
	fmt.Println(countPrimesV2(5))
	fmt.Println(countPrimesV2(10))
}
