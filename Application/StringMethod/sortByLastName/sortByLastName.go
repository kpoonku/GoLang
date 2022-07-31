package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) String() string {
	return fmt.Sprintf("%s %s %d", p.first, p.last, p.age)
}

type ByLastName []person

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Less(i, j int) bool { return a[i].last < a[j].last }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	person := []person{{"Sabareesh", "muthiah Venkat", 6},
		{"Hareni", "mangai Venkat", 3},
		{"Kunthi", "poonkundran", 35},
		{"Venkat", "muthiah", 39}}

	fmt.Println(person)
	sort.Sort(ByLastName(person))
	fmt.Println(person)

}
