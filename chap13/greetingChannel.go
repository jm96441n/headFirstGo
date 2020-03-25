package main

import "fmt"

func greeting(myChannel chan string) {
        myChannel <- "hi"
}

func main() {
        myChannel := make(chan string)
        go greeting(myChannel)
        receivedValue := <-myChannel
        fmt.Println(receivedValue)
}
