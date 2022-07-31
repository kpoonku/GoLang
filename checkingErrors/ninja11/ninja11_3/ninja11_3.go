package main

import (
	"fmt"
	"log"
)

type customErr struct {
	err string
}

func (c customErr) Error() string {
	return fmt.Sprintf("Invalid Number : %v", c.err)
}

func main() {
	r, err := isPositiveNumber(-10)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(r)
	}
}
