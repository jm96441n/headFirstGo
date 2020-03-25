package main

import "fmt"

func abc(channel chan string) {
        channel <- "a"
        channel <- "b"
        channel <- "c"
}

func def(channel chan string) {
        channel <- "d"
        channel <- "e"
        channel <- "f"
}

func main() {
        channel11 := make(chan string)
        channel12 := make(chan string)
        go abc(channel11)
        go def(channel12)
        fmt.Println(<-channel11)
        fmt.Println(<-channel12)
        fmt.Println(<-channel11)
        fmt.Println(<-channel12)
        fmt.Println(<-channel11)
        fmt.Println(<-channel12)
}
