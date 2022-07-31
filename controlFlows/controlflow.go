package main

import "fmt"

func main() {
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	// x := 1
	// for x < 6 {
	// 	fmt.Print(x, "\t")
	// 	x++
	// }
	// fmt.Println()

	// x = 1
	// for {
	// 	if x == 10 {
	// 		break
	// 	}
	// 	fmt.Print(x, "\t")
	// 	x++
	// }

	// fmt.Println("Continue... ")

	// x = 1
	// for {
	// 	x++
	// 	if x >= 10 {
	// 		break
	// 	}
	// 	if x%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Print(x, "\t")
	// }

	//asciiNum := 33 // Uppercase A

	// fmt.Println("Ascii Printing")
	// for asciiNum < 122 {
	// 	character := string(asciiNum)
	// 	fmt.Println(character)
	// 	asciiNum++
	// }
	// fmt.Println()

	// if x := 52; x == 42 {
	// 	fmt.Println("Equals 42")
	// } else if x == 52 {
	// 	fmt.Println(" Equals 52")
	// } else {
	// 	fmt.Println("Not equals 42 or 52")
	// }

	// for i := 3; i < 100; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	switch switchValue := 50; switchValue {
	case 42:
		fmt.Println("42")
	case 43:
		fmt.Println("43")
		fallthrough
	case 44:
		fmt.Println("44")
	case 45:
		fmt.Println("45")
		fallthrough
	case 46, 47, 48, 49:
		fmt.Println("46")
		fallthrough
	default:
		fmt.Println("NA")
	}
}
