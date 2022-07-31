package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, " and I am secret agent")
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " and I can speak")
}

func bar(h human) {
	fmt.Println("I am human , and I can speak")
	fmt.Println(h)
	// Assertion
	switch h.(type) {
	case person:
		fmt.Println("I was passed into Bar,", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{first: "James",
			last: "Bond",
			age:  35,
		},
		ltk: true,
	}

	person := person{
		first: "Will",
		last:  "Smith",
		age:   45,
	}

	fmt.Println(sa1)
	sa1.speak()
	human.speak(sa1)
	human.speak(person)
	bar(sa1)
	bar(person)
}
