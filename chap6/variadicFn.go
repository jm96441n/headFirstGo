package main

import "fmt"

func main() {
        severalInts(1)
        severalInts(1, 2, 3)
        mix(1, true, "a", "b")
        mix(2, false, "a", "b", "c", "d")
}

func twoInts(first int, second int) {
        fmt.Println(first, second)
}

func severalInts(numbers ...int) {
        fmt.Println(numbers)
}

func mix(num int, flag bool, strings ...string) {
        fmt.Println(num, flag, strings)
}
