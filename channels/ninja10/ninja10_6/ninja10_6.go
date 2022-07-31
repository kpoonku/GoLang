package main

import "fmt"

func main() {

	receiveOnlyChannel := gen()

	receive(receiveOnlyChannel)
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(receiveOnlyChannel <-chan int) {
	for v := range receiveOnlyChannel {
		fmt.Println(v)
	}
}
