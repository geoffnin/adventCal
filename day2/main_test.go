package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestExample1RowInputHighLow(t *testing.T) {
	table := []struct {
		input string
		high  int
		low   int
	}{
		{"5 1 9 5", 9, 1},
		{"7 5 3", 7, 3},
		{"2 4 6 8", 8, 2},
		{`414	382	1515	319	83	1327	116	391	101	749	1388	1046	1427	105	1341	1590`, 1590, 83},
		{`960	930	192	147	932	621	1139	198	865	820	597	165	232	417	19	183`, 1139, 19},
	}

	for _, test := range table {
		normalized := normalize(test.input)
		row := rowFromString(normalized)
		high, low := getHighLow(row)
		if high != test.high || low != test.low {
			t.Errorf("gave %v, got %v %v expected %v %v", test.input, high, low, test.high, test.low)
		}
	}
}

func TestExample1Checksum(t *testing.T) {
	table := []struct {
		input    string
		checksum int
	}{
		{`5 1 9 5
7 5 3
2 4 6 8`, 18},
	}

	for _, test := range table {
		checksum := checksum1(bufio.NewScanner(strings.NewReader(test.input)))

		if checksum != test.checksum {
			t.Errorf("gave %v expected %v got %v", test.input, test.checksum, checksum)
		}

	}
}

func TestNormalizeInput(t *testing.T) {
	output := normalize(`	414	382	1515	319	83	1327	116	391	101`)
	if output != `414 382 1515 319 83 1327 116 391 101` {
		t.Errorf("did not normalize correctly '%v' ", output)
	}

}

func TestEvenlyDivisableNumbers(t *testing.T) {
	table := []struct {
		input     []int
		high, low int
	}{
		{[]int{5, 9, 2, 8}, 8, 2},
		{[]int{9, 4, 7, 3}, 9, 3},
		{[]int{3, 8, 6, 5}, 6, 3},
	}

	for _, test := range table {
		high, low := getEvenlyDivisable(test.input)

		if high != test.high || low != test.low {
			t.Errorf("%v : Expected (%v,%v) Got (%v,%v)", test.input, test.high, test.low, high, low)
		}
	}
}
