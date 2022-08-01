package acdc

import (
	"fmt"
	"testing"
)

func ExampleSum() {
	fmt.Println(Sum(2, 3))
	// Output:
	// 5
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(i, i)
	}
}
