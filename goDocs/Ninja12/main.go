package main

import (
	"fmt"
)

type DoberMan struct {
	name string
	age  int
}

func main() {
	fido := DoberMan{
		name: "Fido",
		age:  10,
	}
	fmt.Println(fido)
}
