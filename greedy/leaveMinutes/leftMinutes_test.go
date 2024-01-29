package leetdoce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalMinutes(t *testing.T) {
	assert.Equal(t, 2, calMinutes(1, 2))
}
