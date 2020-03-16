package main

import "fmt"

func main() {
        notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
        index := 1
        fmt.Println(index, notes[index])
        fmt.Println("Printing specific indexes")
        fmt.Println(index, notes[index])
        index = 3
        fmt.Println("LOOP TIME")
        fmt.Println(len(notes))
        for i := 0; i < len(notes); i++ {
                fmt.Println(i, notes[i])
        }

        fmt.Println("FOR RANGE LOOOOOP")
        for index, value := range notes {
                fmt.Println(index, value)
        }

        fmt.Println("NO INDEXES HERE")
        for _, value := range notes {
                fmt.Println(value)
        }

        fmt.Println("NO VALUES HERE")
        for index, _ := range notes {
                fmt.Println(index)
        }

}
