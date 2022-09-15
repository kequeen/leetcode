package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	assert.Equal(t, wordBreak(s, wordDict), true)
}
