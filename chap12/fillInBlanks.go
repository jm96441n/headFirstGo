package main

import "fmt"

func snack() {
        defer fmt.Println("Closing Regfrigerator")
        fmt.Println("Opening Regfrigerator")
        panic("Regfrigerator empty!")
}

func main() {
        snack()
}
