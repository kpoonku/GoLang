package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println("I am human")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	person := person{"Sabareesh", "Venkat"}
	saySomething(&person)
	//saySomething(person)
}
