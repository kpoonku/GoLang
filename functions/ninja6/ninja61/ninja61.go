package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceIntTotal := foo(sliceInt...)
	fmt.Println(sliceIntTotal)
	sliceIntTotal = bar(sliceInt)
	fmt.Println(sliceIntTotal)

}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("sum : ", sum)
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
