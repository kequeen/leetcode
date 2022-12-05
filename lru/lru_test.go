package leetcode

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	capacity := 2
	key := 1
	value := 1
	obj := Constructor(capacity)
	param_1 := obj.Get(key)
	fmt.Println(param_1)
	obj.Put(key, value)
	fmt.Println(obj.Get(key))
}
