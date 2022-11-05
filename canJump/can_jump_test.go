package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	arr := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(arr))
	assert.Equal(t, canJumpV2(arr), true)

	arr2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(arr2))
	assert.Equal(t, canJumpV2(arr2), false)
}
