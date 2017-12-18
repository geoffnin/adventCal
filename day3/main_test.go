package main

import "testing"

func TestExampleCases(t *testing.T) {
	table := []struct {
		input    int
		expected int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 1},
		{7, 2},
		{8, 1},
		{9, 2},
		{10, 3},
		{11, 2},
		{12, 3},
		{13, 4},
		{14, 3},
		{15, 2},
		{23, 2},
		{1024, 31},
	}

	for _, test := range table {
		dist := distance(test.input)
		if dist != test.expected {
			t.Errorf("Got %v for input %v. expected %v", dist, test.input, test.expected)
		}
	}

}
