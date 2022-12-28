package leetcode

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	s1 := "ADOBECODEBANC"
	t1 := "ABC"
	fmt.Println(minWindow(s1, t1))
	fmt.Println(minWindowV2(s1, t1))

	s2 := "a"
	t2 := "b"
	fmt.Println(minWindow(s2, t2))
	fmt.Println(minWindowV2(s2, t2))

	s3 := "ab"
	t3 := "a"
	fmt.Println(minWindow(s3, t3))
	fmt.Println(minWindowV2(s3, t3))

	s4 := "ab"
	t4 := "b"
	fmt.Println(minWindow(s4, t4))
	fmt.Println(minWindowV2(s4, t4))

}
