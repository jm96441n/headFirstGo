package main

import (
        "fmt"
        "math"
)

func squareRoot (number float64) (float64, error) {
        if number < 0 {
                return 0, fmt.Errorf("can't get the square root of a negative number")
        }
        return math.Sqrt(number), nil
}

func main () {
        root, err := squareRoot(-9.3)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Printf("%0.3f\n", root)
        }

        root, err = squareRoot(125)
        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Printf("%0.3f\n", root)
        }

}
