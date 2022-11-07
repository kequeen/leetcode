package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	s := "III"
	fmt.Println(romanToInt(s))

	assert.Equal(t, romanToInt("IV"), 4)

	assert.Equal(t, romanToInt("IX"), 9)
}
