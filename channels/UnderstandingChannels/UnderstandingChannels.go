package main

import "fmt"

func main() {
	/*
		Not Successful
		c := make(chan int)
		c <- 42
		fmt.Println(<-c)
	*/
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	// This works too
	c1 := make(chan int, 1)
	c1 <- 43
	fmt.Println(<-c1)

	// This does not work
	/*c2 := make(chan int, 1)
	c2 <- 43
	c2 <- 44
	fmt.Println(<-c2)*/

	c2 := make(chan int, 2)
	c2 <- 43
	c2 <- 44
	fmt.Println(<-c2)
}
