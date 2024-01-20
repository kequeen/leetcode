package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPrimes(t *testing.T) {
	assert.Equal(t, countPrimes(4), 2)
	assert.Equal(t, countPrimes(5), 2)

	assert.Equal(t, countPrimesV2(4), 2)
	assert.Equal(t, countPrimesV2(5), 2)
}
