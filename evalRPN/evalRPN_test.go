package leetcode

import (
	"fmt"
	"testing"
)

func TestEvalPrn(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
	fmt.Println(evalRPNV2(tokens))

	tokens2 := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens2))
	fmt.Println(evalRPNV2(tokens2))

}
