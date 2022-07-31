package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss MoneyPenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["Barnables"])
	v, ok := m["Barnable"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnables"]; ok {
		fmt.Println("Value : ", v)
	}
	if v, ok := m["James"]; ok {
		fmt.Println("value : ", v)
	}
	delete(m, "James")
	fmt.Println(m)

	//lookup
	v, ok = m["James"]
	fmt.Println(v, ok)

}
