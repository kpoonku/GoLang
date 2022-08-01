package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", sumNumbers(2, 3))
	fmt.Println("2 + 5 =", sumNumbers(2, 5))
	fmt.Println("6 + 5 =", sumNumbers(6, 5))
}

func sumNumbers(i ...int) int {
	result := 0
	for _, v := range i {
		result += v
	}
	return result
}
