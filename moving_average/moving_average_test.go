package leetcode

import (
	"fmt"
	"testing"
)

func TestMovingAverageTest(t *testing.T) {
	obj := Constructor(3)
	param_1 := obj.Next(10)

	fmt.Println(param_1)
	param_2 := obj.Next(2)
	fmt.Println(param_2)

	param_3 := obj.Next(2)
	fmt.Println(param_3)

	param_4 := obj.Next(2)
	fmt.Println(param_4)

}
