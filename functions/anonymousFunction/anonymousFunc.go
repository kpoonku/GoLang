package main

import "fmt"

func foo() {
	fmt.Println("I am foo")
}

func main() {
	foo()
	func() {
		fmt.Println("Anonymous Function No arguments")
	}()

	func(x int) {
		fmt.Println("Anonymous function int arguments ", x)
	}(43)
}
