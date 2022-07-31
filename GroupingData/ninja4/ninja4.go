package main

import "fmt"

func main() {
	var arrayInt [5]int
	for i := 0; i < len(arrayInt); i++ {
		arrayInt[i] = i + 10
	}
	for b, i := range arrayInt {
		fmt.Println(b, i)
	}
	fmt.Printf("%T\n", arrayInt)

	sliceInt := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

	for b, i := range sliceInt {
		fmt.Println(b, i)
	}
	fmt.Printf("%T\n", sliceInt)

	sliceInt1 := sliceInt[2:3]
	fmt.Println(sliceInt1)
	sliceInt2 := sliceInt[3:8]
	fmt.Println(sliceInt2)
	sliceInt3 := sliceInt[4:9]
	fmt.Println(sliceInt3)
	sliceInt4 := sliceInt[6:]
	fmt.Println(sliceInt4)
	sliceInt5 := sliceInt[5:9]
	fmt.Println(sliceInt5)

	sliceInt1 = append(sliceInt[:4], sliceInt[6:]...)
	fmt.Println("Delete from Slice : ", sliceInt1)

	USAStates := make([]string, 0, 50)
	USAStates = append(USAStates, []string{`Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
	Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
	Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
	Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
		` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
	Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}...)

	for i := 0; i < len(USAStates); i++ {
		fmt.Println(i, "\t", USAStates[i])
	}
	fmt.Println(len(USAStates), cap(USAStates))

	slices := [][]string{{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println("slices : ", slices)

	// to throw away index
	for _, v := range slices {
		fmt.Println(v)
		for _, u := range v {
			fmt.Println(u)
		}
	}

	//map of String and slice of string
	handsOnEight := map[string][]string{
		"bond_james":      []string{"Shaken, not storred", "MArtinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "ice cream", "Sunsets"},
	}

	handsOnEight["Hareni_Mangai"] = []string{"Swimming", "Splash Pad", "Reading"}
	for key, element := range handsOnEight {
		fmt.Println("Key : ", key)
		for i, val := range element {
			fmt.Println(i, "\t", val)
		}
	}

	delete(handsOnEight, "no_dr")

	fmt.Printf("\nAfter Deletion ... \n\n\n")

	for key, element := range handsOnEight {
		fmt.Println("Key : ", key)
		for i, val := range element {
			fmt.Println(i, "\t", val)
		}
	}
}
