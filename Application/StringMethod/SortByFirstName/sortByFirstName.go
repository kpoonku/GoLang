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

type ByFirstName []person

func (a ByFirstName) Len() int           { return len(a) }
func (a ByFirstName) Less(i, j int) bool { return a[i].first < a[j].first }
func (a ByFirstName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	person := []person{{"Sabareesh", "Venkat", 6},
		{"Hareni", "Venkat", 3},
		{"Kunthi", "Venkat", 35},
		{"Venkat", "Muthiah", 39}}

	fmt.Println(person)
	sort.Sort(ByFirstName(person))
	fmt.Println(person)

}
