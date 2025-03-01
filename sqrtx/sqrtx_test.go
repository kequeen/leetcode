package sqrtx

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x      int
		expect int
	}{
		{4, 2},
		{8, 2},
		{9, 3},
		{16, 4},
		{25, 5},
		{36, 6},
	}
	for _, tt := range tests {
		if got := mySqrt(tt.x); got != tt.expect {
			t.Errorf("mySqrt(%d) = %d, want %d", tt.x, got, tt.expect)
		}
	}
}
