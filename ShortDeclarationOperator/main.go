package main

import "fmt"

// var declaration
var y = 30
var zy int
var s string

var xxx int = 42
var yyy string = "James Bond"
var zzz bool = true

type Jasmine string

var jasmine Jasmine = "Hellow"

type CountFlower int

var cf CountFlower = 10

func main() {
	// Short Declaration operator
	x := 20
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 70
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
	fmt.Println(zy)
	fmt.Println(s)
	zy = 50
	fmt.Println(y, zy)
	ys := "string type now"

	fmt.Printf("%T\n", ys)
	fmt.Printf("%T\n", s)
	xx := 42
	yy := "James Bond"
	uu := true

	cf = CountFlower(zy)
	fmt.Printf("cf Type : %T\n", cf)
	fmt.Println(cf)

	fmt.Println(xx)
	fmt.Println(yy)
	fmt.Println(uu)
	fmt.Println(xx, yy, uu)
	fmt.Println(xxx)
	fmt.Println(yyy)
	fmt.Println(zzz)
	s := fmt.Sprintf("%v\t%v\t%v", xxx, yyy, zzz)
	fmt.Println(s)
	fmt.Printf("T1 Type : %T\n", jasmine)
	fmt.Println(jasmine)
}
