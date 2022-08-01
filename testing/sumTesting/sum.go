package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", sumNumbers(2, 3))
	fmt.Println("2 + 3.4 =", sumNumbers(2, 3.4))
	fmt.Println("6.2 + 5.6 =", sumNumbers(6.2, 5.6))
}

func sumNumbers(f1, f2 float64) float64 {
	return f1 + f2 + 1
}
