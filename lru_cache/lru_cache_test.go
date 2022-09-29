package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	capacity := 2
	key := 1
	value := 1
	obj := Constructor(capacity)
	param_1 := obj.Get(key)
	fmt.Println(param_1)
	obj.Put(key, value)
	fmt.Println(obj.Get(key))

	key2 := 2
	value2 := 2
	obj.Put(key2, value2)
	param_2 := obj.Get(key2)
	fmt.Println(param_2)

	fmt.Println(obj.Get(key))

	key3 := 3
	value3 := 3
	obj.Put(key3, value3)
	fmt.Println(obj.Get(key3))
	fmt.Println(obj.Get(key2))
}

func TestLRUCache2(t *testing.T) {
	//["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	//[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]

	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	fmt.Println(obj.arr)
	obj.Get(4)
	fmt.Println(obj.arr)
	obj.Get(3)
	fmt.Println(obj.arr)
	obj.Get(2)
	fmt.Println(obj.arr)
	obj.Get(1)
	fmt.Println(obj.arr)

}
