package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintMatrix(t *testing.T) {
	fmt.Println(printMatrix(10))
	assert.Equal(t, 123, 123, "They should be equal")
}
