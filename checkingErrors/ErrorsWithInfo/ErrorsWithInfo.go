package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative numbers")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		//panic("panicking")
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//	return 0, errors.New("norgate math: square root of negative numbers")
		return 0, ErrNorgateMath
	}
	return 42, nil
}
