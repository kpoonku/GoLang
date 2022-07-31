package main

import (
	"fmt"
	"runtime"
)

func main() {

	receiveOnlyChannel := gen()

	receive(receiveOnlyChannel)
}

func gen() <-chan int {
	c := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	return c
}

func receive(receiveOnlyChannel <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-receiveOnlyChannel)
	}
}
