package main

import (
	"fmt"
	"testing"
)

func TestSumNumbers(t *testing.T) {
	result := sumNumbers(2, 3)
	if result != 5 {
		t.Error("Expected", 5, "Got", result)
	} else {
		fmt.Println("sucess ")
	}
}
