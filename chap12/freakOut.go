package main

import "fmt"

func calmDown() {
        p := recover()
        err, ok := p.(error)
        if ok {
                fmt.Println(err.Error())
        }
}

func freakOut() {
        defer calmDown()
        err := fmt.Errorf("oh my, there's been an error")
        panic(err)
        fmt.Println("I won't be run")
}

func main() {
        freakOut()
        fmt.Println("Exiting normally")
}
