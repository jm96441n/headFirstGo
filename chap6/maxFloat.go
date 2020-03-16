package main

import (
        "fmt"
        "math"
)

func max(numbers ...float64) float64 {
        max := math.Inf(-1) // the lowest infinite value
        for _, number := range numbers {
                if number > max {
                        max = number
                }
        }
        return max

}

func main() {
        fmt.Println(max(71.8, 56.2, 89.5))
        fmt.Println(max(90.7, 89.7, 98.5, 92.3))
}
