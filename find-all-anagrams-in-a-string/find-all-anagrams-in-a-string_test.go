package leetcode

import "testing"

func TestFindAnagrams(t *testing.T) {
	s := "aa"
	p := "bb"
	if len(findAnagrams(s, p)) != 0 {
		t.Fatal("findAnagrams error")
	}
}
