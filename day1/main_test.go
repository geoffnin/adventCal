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
		response := sumPart1(test.input)
		if test.expected != response {
			t.Errorf("Gave %v, got %v but expected %v", test.input, response, test.expected)
		}
	}
}

func TestPart2ExamplesReturnCorrectValue(t *testing.T) {
	table := []struct {
		input    string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, test := range table {
		response := sumPart2(test.input)
		if response != test.expected {
			t.Errorf("Gave %v, got %v but expected %v", test.input, response, test.expected)
		}
	}
}
