package main

import "fmt"

func main() {
        slice := []string{"a", "b"}
        fmt.Println(slice, len(slice))

        slice = append(slice, "c")
        fmt.Println(slice, len(slice))

        slice = append(slice, "d", "e")
        fmt.Println(slice, len(slice))

        var slice2 []string
        if len(slice2) == 0 {
                slice2 = append(slice2, "the first item")
        }

        fmt.Printf("%#v\n", slice)

}
