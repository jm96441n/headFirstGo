package main

import "fmt"

func main() {
        var notes []string
        // make instantiates the slice with the second arg being the size
        notes = make([]string, 7)
        notes[0] = "do"
        notes[1] = "re"
        notes[2] = "mi"
        fmt.Println(notes[0])
        fmt.Println(notes[1])

        primes := make([]int, 5)
        primes[0] = 2
        primes[1] = 3
        fmt.Println(primes[0])

        letters := []string{"a", "b", "c"}
        for i := 0; i < len(letters); i++ {
                fmt.Println(letters[i])
        }

        for _, letter := range letters {
                fmt.Println(letter)
        }

        // create a slice from an array
        underlyingArray := [5]string{"a", "b", "c", "d", "e"}
        slice1 := underlyingArray[0:3]
        // don't need to indicate the start index if it's 0
        slice2 := underlyingArray[:3]
        // omitting the end index will grab everything to the end of the array
        slice3 := underlyingArray[2:]
        fmt.Println(slice1)
        fmt.Println(slice2)
        fmt.Println(slice3)

}
