package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(" I am ", p.first, " ", p.last)
}

func main() {
	person := person{
		first: "James",
		last:  "Bond",
		age:   45,
	}
	person.speak()
}
