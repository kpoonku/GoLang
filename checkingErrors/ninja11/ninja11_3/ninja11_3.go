package main

import (
	"errors"
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

func isPositiveNumber(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("negative number")
	}
	return i, nil
}
