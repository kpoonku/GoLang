package main

import "fmt"

func main() {
	iFunc := foo()
	fmt.Println(iFunc())
	fmt.Println(foo()())
	defer bar()
	bar1()
}

func foo() func() int {
	var i int
	fmt.Println("foo")
	return func() int {
		i++
		return i
	}
}

func bar() {
	fmt.Println("Bar defered function")
}

func bar1() {
	fmt.Println("Bar1 function")
}
