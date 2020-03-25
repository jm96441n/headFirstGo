package main

import "fmt"

func one() {
        defer fmt.Println("deferred from one()")
        two()
}

func two() {
        defer fmt.Println("deferred from two()")
        three()
}

func three() {

        panic("sugar we're going down swinging")
}

func main() {
        one()
}
