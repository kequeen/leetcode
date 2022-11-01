package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	s1 := "hello world"
	assert.Equal(t, lengthOfLastWord(s1), 5)

	s2 := "   fly me   to   the moon  "
	assert.Equal(t, lengthOfLastWord(s2), 4)

	s3 := "luffy is still joyboy"
	assert.Equal(t, lengthOfLastWord(s3), 6)
}
