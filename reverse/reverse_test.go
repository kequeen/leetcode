package leetcode

import "testing"

func TestReverse(t *testing.T) {
	num := 321
	if reverse(num) != 123 {
		t.Fatal("不符合预期")
	}

}
