package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	fmt.Println(longestValidParentheses(s))

	s1 := ")()())"
	fmt.Println(longestValidParentheses(s1))

	s2 := ""
	fmt.Println(longestValidParentheses(s2))
}
