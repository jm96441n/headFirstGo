package main

import "fmt"

type Number int

func (n *Number) double() {
	*n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("Original value:", number)
	number.double()
	fmt.Println("New Value:", number)
}
