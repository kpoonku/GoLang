package main

import (
	"fmt"
	"runtime"
)

func maigop() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
