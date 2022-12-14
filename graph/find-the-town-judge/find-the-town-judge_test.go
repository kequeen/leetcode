package leetcode

import (
	"fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {

	n := 2
	trust := [][]int{{1, 2}}
	fmt.Println(findJudge(n, trust))

	n2 := 3
	trust2 := [][]int{{1, 3}, {2, 3}}
	fmt.Println(findJudge(n2, trust2))

	n3 := 3
	trust3 := [][]int{{1, 3}, {2, 3}, {3, 1}}
	fmt.Println(findJudge(n3, trust3))

}
