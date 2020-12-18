package main

import (
	"fmt"
	"sync"
	"time"
)

var count int32
var mutex = &sync.Mutex{}

// use -race to use race detector
func main() {
	fmt.Println("Race condition")
	for i := 0; i < 100; i++ {
		go func(n int) {
			mutex.Lock()
			defer mutex.Unlock()
			count++
			fmt.Printf("i: %d count: %d\n", n, count)
		}(i)

	}
	time.Sleep(2 * time.Second)
}
