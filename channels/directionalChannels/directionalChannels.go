package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 23
	ch <- 32

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("----------")
	fmt.Println("ch : %T\n", ch)

	c := make(chan<- int, 2) // send
	c <- 42
	c <- 43

	// send only channel - cannot receive
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	fmt.Println("----------")
	fmt.Printf("c : %T\n", c)

	c1 := make(<-chan int, 2) // receive
	// receive only channel, cannot send
	// c1 <- 42
	// c1 <- 43

	// send only channel - cannot receive
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	fmt.Printf("----------")
	fmt.Println("c1: %T\n", c1)
	c = ch
	c1 = ch
	fmt.Println(<-c1)

	// does not work
	ch = c1
}
