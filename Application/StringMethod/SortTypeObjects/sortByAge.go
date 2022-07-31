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

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	person := []person{{"Sabareesh", "Venkat", 6},
		{"Hareni", "Venkat", 3},
		{"Kunthi", "Venkat", 35},
		{"Venkat", "Muthiah", 39}}

	fmt.Println(person)
	sort.Sort(ByAge(person))
	fmt.Println(person)

}
