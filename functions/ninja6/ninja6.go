package main

import "fmt"

func main() {
	fmt.Println(foo())
	x, y := bar()
	fmt.Println(x, y)
}

func foo() int {
	return 60
}

func bar() (int, string) {
	return 45, "sabareesh"
}
