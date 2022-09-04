package leetcode

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {

	var mutex sync.RWMutex
	arr := []int{1, 2, 3}
	go func() {
		fmt.Println("Try to lock writing operation.")
		mutex.Lock()
		fmt.Println("Writing operation is locked.")

		arr = append(arr, 4)

		fmt.Println("Try to unlock writing operation.")
		mutex.Unlock()
		fmt.Println("Writing operation is unlocked.")
	}()

	go func() {
		fmt.Println("Try to lock reading operation.")
		mutex.RLock()
		fmt.Println("The reading operation is locked.")

		fmt.Println("The len of arr is : ", len(arr))

		fmt.Println("Try to unlock reading operation.")
		mutex.RUnlock()
		fmt.Println("The reading operation is unlocked.")
	}()

	time.Sleep(time.Second * 2)
}

func TestBufio(t *testing.T) {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!
	a := []byte{123}
	b := []byte{2}
	fmt.Println(bytes.Equal(a, b))
}
