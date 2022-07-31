package main

import (
	"FirstCode/goDocs/Ninja12/dog"
	"fmt"
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
