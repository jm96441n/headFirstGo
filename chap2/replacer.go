package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks"
	fmt.Println(broken)
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
