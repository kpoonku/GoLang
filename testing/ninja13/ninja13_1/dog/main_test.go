package dog

import (
	"fmt"
	"testing"
)

/*
// Years converts human years to dog years.
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

*/
func TestYearsTwo(t *testing.T) {
	result := YearsTwo(1)
	if result != 7 {
		t.Error("Expected : ", 8, "Actual : ", result)
	}
}

func ExampleYearsTwo() {
	n := 2
	fmt.Println(YearsTwo(n))
	// Output:
	// 14
}

func BenchmarkYearsTwo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}

func TestYears(t *testing.T) {
	result := Years(1)
	if result != 7 {
		t.Error("Expected : ", 7, "Actual : ", result)
	}
}

func ExampleYears() {
	n := 2
	fmt.Println(Years(n))
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}
