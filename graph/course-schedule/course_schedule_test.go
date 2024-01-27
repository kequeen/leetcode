package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	assert.Equal(t, canFinish(numCourses, prerequisites), true)
	assert.Equal(t, canFinish(numCourses, prerequisites), true)

	numCourses2 := 2
	prerequisites2 := [][]int{{1, 0}, {0, 1}}
	assert.Equal(t, canFinish(numCourses2, prerequisites2), false)
	assert.Equal(t, canFinishV2(numCourses2, prerequisites2), false)
}
