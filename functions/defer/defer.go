package main

import "fmt"

func foo(i int) {
	fmt.Println("I am foo , i: ", i)
}

func bar() {
	fmt.Println("I am bar...")
}

func main() {
	defer foo(1)
	bar()
}
