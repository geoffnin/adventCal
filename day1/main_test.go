package main

import "testing"

func TestExamplesReturnCorrectValue(t *testing.T) {
	table := []struct {
		input    string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, test := range table {
		response := sumFromString(test.input)
		if test.expected != response {
			t.Errorf("Gave %v, got %v expected %v", test.input, response, test.expected)
		}
	}
}
