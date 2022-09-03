package leetcode

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	obj := Constructor()

	obj.AddNum(6)
	obj.AddNum(10)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	if obj.FindMedian() != 6 {
		t.Error("fail")
	}
	obj.AddNum(6)
	obj.AddNum(5)

	obj.AddNum(0)

	fmt.Println(obj.FindMedian())
	obj.AddNum(6)
	obj.AddNum(3)
	obj.AddNum(1)

	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	obj.AddNum(0)

	obj.AddNum(16)
	obj.AddNum(8)
	obj.AddNum(10)
	obj.AddNum(12)
	obj.AddNum(0)
	obj.AddNum(16)
	fmt.Println(obj.FindMedian())

	// newArr := []int{1, 3, 4, 5, 6, 2, 3}
	// fmt.Println(newArr[0:3])
	// fmt.Println(newArr[3:])
	// insertValue := 100
	// newArr = append(append(newArr[0:3], insertValue), newArr[3:]...)
	// fmt.Println(newArr)
}
