package main

import "fmt"

func main() {
        var numbers [3]int
        numbers[0] = 42
        numbers[2] = 108
        var letters = [3]string{"a", "b", "c"}

        // print 42
        fmt.Println(numbers[0])

        // print 0
        fmt.Println(numbers[1])
       
        // print 108
        fmt.Println(numbers[2])
       
        // print c
        fmt.Println(letters[2])
        
        // print a
        fmt.Println(letters[0])

        // print b
        fmt.Println(letters[1])
}
