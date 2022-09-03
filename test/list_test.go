package leetcode

import (
	"container/list"
	"fmt"
	"testing"
)

//测试下 List 的一些功能
func TestList(t *testing.T) {
	var l = list.New()
	e0 := l.PushBack(2)
	l.PushBack("a")
	l.PushFront(5)
	l.InsertBefore(4, e0)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \n", e.Value)
	}

}
