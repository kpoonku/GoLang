package main

import "fmt"

func main() {
	a := 90
	fmt.Println("a : ", a)
	fmt.Println("&a : ", &a)
	fmt.Println("*&a : ", *&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Printf("%T\n", *&a)

	// use pointer to int
	// pass by value in go
	var b *int = &a
	fmt.Println("b *int : ", b)
	// get the value at an address
	fmt.Println("b *int value : ", *b)

	foo(0)
	foo1(0)
}

func foo(y int) {
	x := 0
	bar(x)
	fmt.Println("x :", x)
}

func bar(yi int) {
	fmt.Println("yi : ", yi)
	yi = 44
	fmt.Println("yi : ", yi)
}

func foo1(y int) {
	x := 0
	fmt.Println("x befor : ", &x)
	fmt.Println("x befor : ", x)
	bar1(&x)
	fmt.Println("x after : ", &x)
	fmt.Println("x after : ", x)
}

func bar1(yi *int) {
	fmt.Println("yi befor : ", yi)
	fmt.Println("yi befor : ", *yi)
	*yi = 43
	fmt.Println("yi after : ", yi)
	fmt.Println("yi after : ", *yi)
}
