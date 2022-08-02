package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	mapInt := gen()
	for k, v := range mapInt {
		avgResult := CenteredAvg(v)
		if avgResult != k {
			t.Error("Expected", k, "Actual", avgResult)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output:
	//3
}

func gen() map[float64][]int {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{2, 3, 4, 5, 6, 7}
	c := []int{90, 23, 13, 99, 56, 34, 67}
	d := []int{12, 3, 9, 45, 14, 89, 78}
	e := []int{34, 54, 46, 64, 78, 87}
	f := []int{23, 32, 56, 36, 78, 48}
	mapInt := make(map[float64][]int)
	mapInt[3.5] = a
	mapInt[4.5] = b
	mapInt[54] = c
	mapInt[31.6] = d
	mapInt[60.5] = e
	mapInt[43] = f
	return mapInt
}

func BenchmarkCenteredAvg(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{i + 6, i, i + 8, i + 2, i + 4, i + 3})
	}
}
