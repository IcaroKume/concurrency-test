package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count int32

// use -race to use race detector
func main() {
	fmt.Println("Race condition")
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
			// fmt.Printf("%d\n", count)
			fmt.Printf("%d\n", atomic.LoadInt32(&count))
		}()

	}
	time.Sleep(2 * time.Second)
}
