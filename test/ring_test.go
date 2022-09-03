package leetcode

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	r := ring.New(5)
	fmt.Printf("ring %v", *r)
	a := make([]int, 2)
	fmt.Println(a)
	fmt.Println(&a[0])
	fmt.Println(&a[1])

	//验证map是可以达到传引用的效果
	k := make(map[int]int)
	k[2] = 4
	k[3] = 5
	fmt.Println(k)
	changeMap(k)
	fmt.Println(k)

	//验证slice
	arr1 := []int{2, 3, 4}
	arr1[2] = 5
	fmt.Println(arr1)
	changeSlice(arr1)
	arr1 = append(arr1, 14)
	fmt.Println(arr1)

	//验证数组
	arr2 := [3]int{2, 4, 5}
	arr2[2] = 1
	fmt.Println(arr2)
	changeArr(arr2)
	fmt.Println(arr2)

}

func changeArr(arr [3]int) {
	arr[1] = 0
}

func changeMap(map1 map[int]int) {
	map1[1] = 2
	map1[2] = 1
}

func changeSlice(arr []int) {
	arr[0] = 5
}
