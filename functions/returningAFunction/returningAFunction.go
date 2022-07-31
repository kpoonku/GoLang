package main

import "fmt"

func main() {
	s := foo()
	fmt.Println(s)
	dfunc := bar()
	fmt.Println("calling returned function : ", dfunc())
	fmt.Printf("%T\n", dfunc)
	fmt.Println(bar()())
}

func foo() string {
	s := "Hello World"
	return s
}

//returning a function
func bar() func() int {
	return func() int {
		return 7898
	}
}
