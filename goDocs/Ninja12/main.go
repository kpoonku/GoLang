package main
import {
	"/github.com/kpoonku/GoLang/tree/master/goDocs/Ninja12/Dog"
}

type DoberMan struct {
	name string
	age int
}

func main() {
	fido := DoberMan{
		name:"Fido",
		age : Dog.Years(10)
	}
	fmt.Println(fido)
}
