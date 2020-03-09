package main

import (
	"fmt"
)

func main() {
	// you can use short variable declaration when delcaring and assigning variables at once using the `:=` assignment
	// operator
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "RiFf RaFf"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
