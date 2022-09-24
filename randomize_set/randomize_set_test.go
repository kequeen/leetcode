package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizeSet(t *testing.T) {
	r := Constructor()
	r.Insert(1)
	r.Insert(3)
	assert.Equal(t, r.Insert(2), true)
	assert.Equal(t, r.Insert(3), false)

	assert.Equal(t, r.Remove(2), true)
	assert.Equal(t, r.Remove(2), false)

	fmt.Println(r.GetRandom())

}
