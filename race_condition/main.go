package main

import (
	"fmt"
	"time"
)

var count []int

// use -race to use race detector
func main() {
	fmt.Println("Race condition")
	for i := 0; i < 100; i++ {
		go func(n int) {
			count = append(count, n)
			fmt.Printf("%d\n", count)
		}(i)

	}
	time.Sleep(2 * time.Second)
}
