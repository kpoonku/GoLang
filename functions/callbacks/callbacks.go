package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	summ := sum(sliceInt...)
	fmt.Println("Sum : ", summ)
	evenTotal := getEvenNumbersTotal(sum, sliceInt...)
	fmt.Println("Even Numbers Total : ", evenTotal)
	oddTotal := getOddNumbersTotal(sum, sliceInt...)
	fmt.Println("Odd Numbers total : ", oddTotal)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func getEvenNumbersTotal(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func getOddNumbersTotal(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
