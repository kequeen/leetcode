package leetcode

import (
	"fmt"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	a1 := "ABCABC"
	b1 := "ABC"
	if gcdOfStrings(a1, b1) != "ABC" {
		t.Error("test error")
	}
	a2 := "ABABAB"
	b2 := "ABAB"
	if gcdOfStrings(a2, b2) != "AB" {
		t.Error("test error")
	}
	a3 := "LEET"
	b3 := "CODE"
	if gcdOfStrings(a3, b3) != "" {
		t.Error("test error")
	}

	a4 := "ABCDEF"
	fmt.Println(a4[4:4])
}
