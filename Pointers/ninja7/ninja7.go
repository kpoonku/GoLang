package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	x := 100
	fmt.Println("Address of x : ", &x)
	person := person{
		first: "James",
		last:  "Bond",
		age:   45,
	}
	fmt.Println("before person : ", person)
	changeMe(&person)
	fmt.Println("after person : ", person)
}

func changeMe(person *person) {
	(*person).first = "Miss"
	(*person).last = "MoneyPenny"
	(*person).age = 69
}
