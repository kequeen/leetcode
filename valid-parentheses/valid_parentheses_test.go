package leetcode

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	s := "{}{}[]"
	fmt.Println(isValid(s))
}
