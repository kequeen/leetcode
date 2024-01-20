package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalPrn(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	assert.Equal(t, evalRPN(tokens), 9)
	assert.Equal(t, evalRPNV2(tokens), 9)

	tokens2 := []string{"4", "13", "5", "/", "+"}
	assert.Equal(t, evalRPN(tokens2), 6)
	assert.Equal(t, evalRPNV2(tokens), 6)
}
