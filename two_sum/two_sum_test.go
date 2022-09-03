package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	//关于二维数组的吐槽
	arr := make([][]int, 5)
	fmt.Println(arr)
	for i := 0; i < 5; i++ {
		arr[i] = make([]int, 5)
	}
	fmt.Println(arr)

	//关于字符串的遍历问题
	str := "al"
	strLen := len(str)
	fmt.Println(strLen)
	for i := 0; i < strLen; i++ {
		fmt.Println(str[i])
	}
	for _, item := range str {
		fmt.Println(item)
	}

	target := 10
	arr2 := []int{2, 3, 4, 1, 4, 5, 6}
	fmt.Println(twoSum(arr2, target))

}
