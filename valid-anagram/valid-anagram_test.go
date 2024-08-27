package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	// 	示例 1:

	// 输入: s = "anagram", t = "nagaram"
	// 输出: true
	// 示例 2:

	// 输入: s = "rat", t = "car"
	// 输出: false

	assert.Equal(t, isAnagram("anagram", "nagaram"), true)
	assert.Equal(t, isAnagram("rat", "car"), false)
}
