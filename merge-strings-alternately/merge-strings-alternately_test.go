package leetcode

import "testing"

func TestMergeAlternately(t *testing.T) {
	a1 := "ab"
	b1 := "pqrs"
	if mergeAlternately(a1, b1) != "apbqrs" {
		t.Error("test error")
	}
	a2 := "abcd"
	b2 := "pq"
	if mergeAlternately(a2, b2) != "apbqcd" {
		t.Error("test error")
	}
}
