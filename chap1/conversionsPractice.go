package main

import "fmt"

func main() {
	length := 1.2
	width := 2
	length = float64(width)
	fmt.Println(length)

	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
}
