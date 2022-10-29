package leetcode

import (
	"fmt"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	s := "234"
	fmt.Println(letterCombinations(s))
	//[adg adh adi aeg aeh aei afg afh afi bdg bdh bdi beg beh bei bfg bfh bfi cdg cdh cdi ceg ceh cei cfg cfh cfi]
	s1 := "23"
	fmt.Println(letterCombinations(s1))
	//[ad ae af bd be bf cd ce cf]
}
