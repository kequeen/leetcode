package leetcode

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {

	s := "10+9/3-1"
	fmt.Println(calculate(s))

	s1 := "12/3+5"
	fmt.Println(calculate(s1))

}
