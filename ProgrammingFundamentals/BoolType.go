package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

var b bool

const (
	a  int     = 42
	bb float64 = 42.78
	c  string  = "James Bond"
)

const (
	aaIota = iota
	bbbIota
	cIota
)

func main() {
	b = true
	fmt.Println(b)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// UTF-8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	stringH := "H"
	fmt.Println(stringH)

	bs = []byte(stringH)
	n := bs[0]
	fmt.Println("Decimal", n)
	fmt.Printf("Type : %T\n", n)
	fmt.Printf("Binary : %b\n", n)
	fmt.Printf("Hex : %x\n", n)
	fmt.Printf("UTF-8 : %#U\n", n)

	fmt.Printf("UTF-8 : %U\n", n)

	fmt.Println(a)
	fmt.Printf("Type : %T\n", a)
	fmt.Println(bb)
	fmt.Printf("Type : %T\n", bb)
	fmt.Println(c)
	fmt.Printf("Type : %T\n", c)

	fmt.Println(aaIota)
	fmt.Printf("Type : %T\n", aaIota)
	fmt.Println(bbbIota)
	fmt.Printf("Type : %T\n", bbbIota)
	fmt.Println(cIota)
	fmt.Printf("Type : %T\n", cIota)

	//Bit Shifting
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	ybitshifting := x << 1
	fmt.Printf("%d\t\t%b\n", ybitshifting, ybitshifting)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	const (
		byt    = iota
		kbIota = 1 << (iota * 10)
		mbIota = 1 << (iota * 10)
		gbIota = 1 << (iota * 10)
	)

	fmt.Printf("%d\t\t\t%b\n", byt, byt)
	fmt.Printf("%d\t\t\t%b\n", kbIota, kbIota)
	fmt.Printf("%d\t\t\t%b\n", mbIota, mbIota)
	fmt.Printf("%d\t\t%b\n", gbIota, gbIota)
}
