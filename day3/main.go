package main

import "fmt"

func wrap(width, high, index int) int {

	return 0
}

func distance(length int) int {
	width := 1
	for high := 1; high <= length; {
		perimeter := (width * 4) - 4
		high += perimeter

		if high >= length {
			return 0
		}

		width += 2
	}

	return 0
}

func main() {
	fmt.Println(distance(265149))
}
