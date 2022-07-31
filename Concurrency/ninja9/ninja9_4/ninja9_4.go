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

	var mutex sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v1 := counter
			v1++
			counter = v1
			mutex.Unlock()
			fmt.Println("counter", counter)
			wg.Done()
		}()
		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
