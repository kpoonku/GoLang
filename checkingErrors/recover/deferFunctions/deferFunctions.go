package main

import "fmt"

func main() {
	foo()
	fmt.Println("-----------------")
	bar()
	fmt.Println("coo : ", coo())
}

func foo() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

func bar() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i, " ")
	}
}

func coo() (i int) {
	defer func() { i++ }()
	return 1
}
