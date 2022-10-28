package leetcode

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	ans := generateParenthesis(3)
	fmt.Println(ans)
	fmt.Println(generateParenthesis(1))
}
