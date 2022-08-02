package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(xi[1], "index position 1")
	fmt.Println(len(xi), "len of xi")
	fmt.Println(xi[len(xi)-1], "value at xi[len(xi)-1] or xi[5]")
	fmt.Println(xi[1:(len(xi) - 1)])
}
