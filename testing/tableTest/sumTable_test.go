package main

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {

	type testData struct {
		data   []int
		answer int
	}

	tests := []testData{
		testData{[]int{21, 21}, 42},
		testData{[]int{2, 2}, 4},
		testData{[]int{1, 1}, 2},
		testData{[]int{0, 0}, 0},
	}

	for _, v := range tests {
		result := sumNumbers(v.data...)
		if result != v.answer {
			t.Error("Expected", v.answer, "Got", result)
		}
	}
}
