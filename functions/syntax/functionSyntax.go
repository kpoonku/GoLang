package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("MoneyPenny")
	fmt.Println(s1)
	x, y := mouse("Ian ", "Fleming")
	fmt.Println(x, y)
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	sliceInt := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	total = sum(sliceInt...)
	fmt.Println("Total Value : ", total)

	xx, yy := 32, 45
	yy, xx = xx, yy

	fmt.Println("print xx, yy : ", xx, yy)

}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}

// variadic parameters
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i,
			"we are now adding", v, "  to the total which is now ", sum)
	}
	fmt.Println("The total is ", sum)
	return sum
}
