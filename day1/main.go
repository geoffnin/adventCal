package main

import (
	"fmt"
	"strconv"
)

type intNode struct {
	value int
	next  *intNode
}

type circularList struct {
	head, tail *intNode
	length     int
}

func (c *circularList) IsEmpty() bool {
	if c.head == nil {
		return true
	}
	return false
}

func (c *circularList) Append(value int) {
	node := intNode{value, nil}
	if c.IsEmpty() {
		c.head = &node
		c.tail = c.head
	} else {
		c.tail.next = &node
	}
	node.next = c.head
	c.length++

	fmt.Println(value, c.head, c.tail, c.String())
}

func (c *circularList) PairsValues() (output [][]int) {
	output = make([][]int, c.length)
	cursor := c.head
	for i := 0; i < c.length; i++ {
		output[i] = []int{cursor.value, cursor.next.value}
		cursor = cursor.next
	}

	return
}

func pairs(input []int) (output [][]int) {
	output = make([][]int, len(input))
	for i := range input {
		if (i + 1) < len(input) {
			output[i] = []int{input[i], input[i+1]}
		} else {
			output[i] = []int{input[i], input[0]}
		}
	}
	return
}

func (c *circularList) String() string {
	output := ""
	cursor := c.head
	for i := 0; i < c.length; i++ {
		output += fmt.Sprintf("%v ", cursor.value)
		cursor = cursor.next
	}
	return output
}

func filter(input [][]int, f func([]int) bool) (output []int) {
	for _, elem := range input {
		if f(elem) {
			output = append(output, elem[0])
		}
	}
	return
}

func sum(input []int) (output int) {
	for _, elem := range input {
		output += elem
	}
	return
}

func sumFromString(input string) (output int) {
	list := make([]int, len(input))

	for i, r := range input {
		value, _ := strconv.Atoi(string(r))
		list[i] = value
	}

	pairs := pairs(list)
	filtered := filter(pairs, func(input []int) bool {
		return input[0] == input[1]
	})
	return sum(filtered)
}

func main() {
	fmt.Println("hello")
}
