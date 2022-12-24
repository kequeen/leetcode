package leetcode

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	s := "aa"
	p := "a"
	fmt.Println(isMatch(s, p))

	s1 := "aa"
	p1 := "a*"
	fmt.Println(isMatch(s1, p1))
}
