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

func isPositiveNumber(inn int) (bool, error) {
	if inn < 0 {
		return false, customErr{"negative number"}
	}
	return true, nil
}
