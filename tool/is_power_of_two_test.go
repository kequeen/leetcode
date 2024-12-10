package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{16, true},
		{32, true},
		{0, false},
		{3, false},
		{5, false},
		{6, false},
		{7, false},
		{9, false},
		{10, false},
	}

	for _, tt := range tests {
		got := IsPowerOfTwo(tt.n)
		assert.Equal(t, tt.want, got, "IsPowerOfTwo(%d) = %t, want %t", tt.n, got, tt.want)
	}
}
