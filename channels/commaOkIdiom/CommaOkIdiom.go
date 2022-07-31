package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)
	fmt.Println("about to exit")

}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value receive from even channel : ", v)
		case v := <-odd:
			fmt.Println("the value receive from even channel : ", v)
		case i, ok := <-quit:
			if ok {
				fmt.Println("frmo comma ok : ", i)
			} else {
				fmt.Println("from comma ok : ", i)
				return
			}
			/*
				// removing unused codes
					case _,_ = <-quit:
						return
					//or
					case <-quit:
						return
			*/
		}
	}
}
