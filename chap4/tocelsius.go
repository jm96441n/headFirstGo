package main

import (
        "fmt"
        "log"
        "github.com/headfirstgo/keyboard"
)

func main () {
        fmt.Print("Enter a temperature in fahrenheit: ")
        fahrenheit, err := keyboard.GetFloat()
        if err != nil {
                log.Fatal(err)
        }
        celsius := (fahrenheit - 32) * (5.0 / 9.0)
        fmt.Printf("%0.2f degress Celsius\n", celsius)
}
