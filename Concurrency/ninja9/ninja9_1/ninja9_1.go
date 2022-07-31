package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Hands-on exercise #1
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/

func main() {
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Main GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("I am main function")
	var wg sync.WaitGroup

	wg.Add(2)
	for i := 0; i < 3; i++ {
		go foo(wg)
		go bar(wg)
	}
	wg.Wait()

	fmt.Println("Main End CPUs : ", runtime.NumCPU())
	fmt.Println("Main End GoRoutines: ", runtime.NumGoroutine())
}

func foo(wg sync.WaitGroup) {
	fmt.Println("Foo Go Routines: ", runtime.NumGoroutine())
	fmt.Println("Foo method")
	wg.Done()
}

func bar(wg sync.WaitGroup) {
	fmt.Println("Bar Go Routines: ", runtime.NumGoroutine())
	fmt.Println("Bar method")
	wg.Done()
}
