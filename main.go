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
	concurrencyV3()

	// arr := [3]int{1, 2, 3}
	// slice := arr[0:1]
	// slice = append(slice, 4)
	// fmt.Println(slice)
	// fmt.Println(arr)

	// slice = append(slice, 5, 6)
	// fmt.Println(slice)
	// fmt.Println(arr)

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
