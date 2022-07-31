package main

import (
	"fmt"
	"time"
)

func main() {
	fainin()
}

func generate(data string) <-chan string {
	stringChannel := make(chan string)
	go func() {
		for {
			stringChannel <- data
			time.Sleep(time.Duration(100 * time.Millisecond))
		}
	}()
	return stringChannel
}

func fainin() {
	c1 := generate("Hello")
	c2 := generate("There")

	fanin := make(chan string)
	go func() {
		for {
			select {
			case s1 := <-c1:
				fanin <- s1
			case s2 := <-c2:
				fanin <- s2
			}
		}
	}()
	go func() {
		for {
			fmt.Println(<-fanin)
		}
	}()
}
