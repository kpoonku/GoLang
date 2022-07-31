package main

import "fmt"

func main() {
	// for i := 1; i <= 10000; i++ {
	// 	fmt.Print(i, " , ")
	// }
	// fmt.Println("")

	// for i := 65; i <= 90; i++ {
	// 	fmt.Printf("%#U\n", i)
	// }

	// bd := 1985
	// for {
	// 	if bd > 2017 {
	// 		break
	// 	}
	// 	fmt.Println(bd)
	// 	bd++
	// }
	// for i := 10; i < 100; i++ {
	// 	fmt.Println(i, " Modulus % 4 : ", i%4)
	// }
	x := 11
	if x == 10 {
		fmt.Println("X is 10")
	} else if x == 11 {
		fmt.Println("x is 11")
	} else {
		fmt.Println("X is 12")
	}

	switch favSport := "surfing"; favSport {
	case "skiing":
		fmt.Println("skiing")
	case "swimming":
		fmt.Println("skiing")
	case "surfing":
		fmt.Println("surfing")
	}

	fmt.Println("true && true : ", true && true)
	fmt.Println("true && false : ", true && false)
	fmt.Println("true || true : ", true || true)
	fmt.Println("true || false : ", true || false)
	fmt.Println("!true : ", !true)
}
