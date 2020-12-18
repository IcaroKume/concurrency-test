package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// waiting until all go routines finish
func main() {
	fmt.Println("Wait group")
	for count := 0; count < 100; count++ {
		wg.Add(1)

		go func(n int) {
			fmt.Printf("%d\n", n)
			wg.Done()
		}(count)
	}
	wg.Wait()
}
