package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	assert.Equal(t, minDistance(word1, word2), 3)

	word3 := "intention"
	word4 := "execution"
	assert.Equal(t, minDistance(word3, word4), 5)

	word5 := "had"
	word6 := "had"
	assert.Equal(t, minDistance(word5, word6), 0)
}
