package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency")
	for count := 0; count < 100; count++ {
		go func(n int) {
			fmt.Printf("%d\n", n)
		}(count)
	}
	time.Sleep(2 * time.Second)
}
