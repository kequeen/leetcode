package main

import (
	"fmt"
	"sync"

	"github.com/panjf2000/ants/v2"
)

// 主要是做golang的各种功能的测试
// 关于channel
func main() {
	// fmt.Println("hello,world")
	// ch := make(chan int)
	// go func() {
	// 	ch <- 3 + 4
	// 	ch <- 10
	// 	ch <- 27
	// }()
	// i := <-ch
	// fmt.Println(i)
	// j := <-ch
	// fmt.Println(j)
	// k := <-ch
	// fmt.Println(k)

	// concurrency()
	// concurrencyV3()

	// arr := [3]int{1, 2, 3}
	// slice := arr[0:1]
	// slice = append(slice, 4)
	// fmt.Println(slice)
	// fmt.Println(arr)

	// slice = append(slice, 5, 6)
	// fmt.Println(slice)
	// fmt.Println(arr)

	// s := []int{1, 1, 1}
	// f(s)
	// fmt.Println(s)
	syncMap()
	// golang := Go{}
	// php := PHP{}

	// sayHello(golang)
	// sayHello(php)
}

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
		i++
	}
	*/

	for i := range s {
		s[i] += 1
	}
}

func syncMap() {
	map1 := make(map[int]int)
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(a int) {
			map1[a] = a
			wg.Done()
		}(i)
	}
	wg.Wait()
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}

func calSlice() {
	s := make([]int, 0)

	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)

		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}

func concurrency() {
	ids := []int{100, 2, 1, 1, 1, 1, 100, 3, 2, 5, 8}
	var wg sync.WaitGroup
	for _, v := range ids {
		syncCalculateSum := func() {
			fmt.Println(v)
		}
		wg.Add(1)
		ants.Submit(syncCalculateSum)
		wg.Done()
	}
	wg.Wait()
	fmt.Printf("finish all tasks.\n")
}

func concurrencyV2() {
	ids := []int{100, 2}
	var wg sync.WaitGroup
	for _, v := range ids {
		syncCalculateSum := func() {
			fmt.Println(v)
			wg.Done()
		}
		wg.Add(1)
		ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("finish all tasks.\n")
}

func concurrencyV3() {
	var wg sync.WaitGroup
	nums := []int{1, 2, 3}

	for _, i := range nums {
		wg.Add(1)
		go func() {
			fmt.Println(i)
		}()
		defer wg.Done()
	}

	wg.Wait()
	fmt.Println("All workers have completed.")
}
