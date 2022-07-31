package main

import "fmt"

// Hello returns a greeting for the named person.
func main() {
	n, err := fmt.Println("Hello World", "26 Bytes")
	fmt.Println(n)
	fmt.Println(err)
}
