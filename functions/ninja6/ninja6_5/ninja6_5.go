package main

import (
	"fmt"
)

type circle struct {
	radius float64
}

type square struct {
	length int
	width  int
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (sq square) area() float64 {
	return float64(sq.length) * float64(sq.width)
}

type shape interface {
	area() float64
}

func main() {
	circle := circle{
		radius: 3,
	}
	square := square{
		length: 3,
		width:  4,
	}
	fmt.Println("circle area : ", circle.area())
	fmt.Println("square area : ", square.area())

	fmt.Println("circle area (interface) : ", shape.area(circle))
	fmt.Println("square area (interface) : ", shape.area(square))

	// anonymous Function
	i := func(i int) int {
		fmt.Println("Anonymous Function ", i)
		return i
	}(66)

	function := func() {
		fmt.Println("Calling function using variable")
	}

	fmt.Println(i)
	function()

	total := callbackCalling(sum, []int{1, 2, 3, 4, 5}...)
	fmt.Println("total : ", total)

	functionA := returnFunction()
	fmt.Println("Call functionA : ", functionA())

	funct := closureFunc()

	fmt.Println("Closure Function : ", funct())

	fmt.Println("Closure Function : ", funct())

	fmt.Println("Closure Function : ", funct())
}

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

//callback - passing a function as an argument to a function
func callbackCalling(funcc func(xi ...int) int, xi ...int) int {
	return funcc(xi...)
}

//returning a function
func returnFunction() func() int {
	return func() int {
		return 98
	}
}

func closureFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
