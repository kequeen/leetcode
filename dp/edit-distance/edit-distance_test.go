package leetcode

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))

	word3 := "intention"
	word4 := "execution"
	fmt.Println(minDistance(word3, word4))
}
