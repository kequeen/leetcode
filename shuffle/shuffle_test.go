package leetcode

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	arr := []int{1, 2, 3, 5}
	obj := Constructor(arr)
	printArr(obj.arr)
	printArr(obj.Shuffle())
	printArr(obj.Shuffle())

}

func printArr(arr []int) {
	for _, value := range arr {
		print(value, " ")
	}
}
