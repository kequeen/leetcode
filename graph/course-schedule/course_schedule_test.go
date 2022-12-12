package leetcode

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites))

	numCourses2 := 2
	prerequisites2 := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses2, prerequisites2))

}
