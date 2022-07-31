package main

import "fmt"

type hotdog int

const (
	a hotdog = 10
	b        = 20
)

const (
	year2018 = 2018 + iota
	year2019 = year2018 + iota
	year2020 = year2019 + iota
	year2021 = year2020 + iota
)

func main() {
	x := 42
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	y := 49
	m := 30
	fmt.Println("x == y", x == y)
	fmt.Println("x <= y", x <= y)
	fmt.Println("x >= y", x >= y)
	fmt.Println("x != y", x != y)
	fmt.Println("x < y", x < y)
	fmt.Println("x > m", x > m)

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

	xShift := 42 << 1
	fmt.Printf("xShift : %d\n", xShift)
	fmt.Printf("xShift : %b\n", xShift)
	fmt.Printf("xShift : %#x\n", xShift)

	yShift := 42 >> 1
	fmt.Printf("yShift : %d\n", yShift)
	fmt.Printf("yShift : %b\n", yShift)
	fmt.Printf("yShift : %#x\n", yShift)

	rawString := `"Here is something fun about 
		learning go language and having 
		family from India home"`
	fmt.Println(rawString)

	fmt.Println("year2018 : ", year2018)
	fmt.Println("year2019 : ", year2019)
	fmt.Println("year2020 : ", year2020)
	fmt.Println("year2021 : ", year2021)

}
