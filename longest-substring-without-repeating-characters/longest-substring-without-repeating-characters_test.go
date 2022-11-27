package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	fmt.Println(lengthOfLongestSubstringV2(s))
	fmt.Println(lengthOfLongestSubstringV3(s))

	s1 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s1))
}
