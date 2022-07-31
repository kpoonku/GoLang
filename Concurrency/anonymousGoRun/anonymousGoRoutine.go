package main

import "fmt"

func doSomething(x int) int {
	//Does something, already built
	return x * 5
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}
