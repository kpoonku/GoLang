package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 10

	var wg sync.WaitGroup
	const gs = 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v1 := counter
			v1++
			counter = v1
			wg.Done()
		}()
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
