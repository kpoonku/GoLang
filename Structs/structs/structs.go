package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "MoneyPenny",
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
