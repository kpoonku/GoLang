package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.GOOS :   ", runtime.GOOS)
	fmt.Println("runtime.GOARCH : ", runtime.GOARCH)
}
