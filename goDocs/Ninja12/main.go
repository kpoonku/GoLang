package main

import (
	"fmt"

	"ninja12.godocs/dog"
)

type DoberMan struct {
	name string
	age  int
}

func main() {
	fido := DoberMan{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}
