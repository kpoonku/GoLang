package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for v := 0; v < len(x); v++ {
		fmt.Println(v, x[v])
	}
	//slice operator
	fmt.Println(x[2:4])
	fmt.Println(x[1:])
	fmt.Println(x[:5])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i, x[i])
	}

	//append
	y := []int{23, 24, 25, 266}
	x = append(x, y...)
	fmt.Println("append : ", x)

	//delete
	z := append(x[:2], x[4:]...)
	fmt.Println("Delete index 2,3 : ", z)

	makeX := make([]int, 10, 20)
	fmt.Println(makeX, len(makeX), cap(makeX))
	makeX[9] = 20
	makeX = append(makeX, y...)
	fmt.Println(makeX, len(makeX), cap(makeX))

	makeX = append(makeX, y...)
	fmt.Println(makeX, len(makeX), cap(makeX))

	makeX = append(makeX, y...)
	fmt.Println(makeX, len(makeX), cap(makeX))

	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "strawberry", "bubblegum"}
	fmt.Println(mp)

	xp := [][]string{{"James", "Bond", "Chocolate", "martini"},
		{"Miss", "MoneyPenny", "strawberry", "bubblegum"}}
	fmt.Println(xp)

	for i, v := range xp {
		fmt.Println(i, v)
		for j, m := range v {
			fmt.Println(j, m)
		}
	}
}

// SLICE - allows to group VALUEs of the same type
// type{values}
