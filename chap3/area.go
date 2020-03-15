package main

import (
        "fmt"
)

var metersPerLiter float64

func paintNeeded(width float64, height float64) (float64, error) {
        if (width < 0) {
                return 0, fmt.Errorf("a width of %0.2f if invalid", width)
        }
        if (height < 0) {
                return 0, fmt.Errorf("a height of %0.2f if invalid", height)
        }
        area := width * height
        return area/10.0, nil
}

func PrintPaintNeeded (area float64) {
        resultString := fmt.Sprintf("%.2f liters needed", area)
        fmt.Println(resultString)
}

func main() {
        var width, height, area float64
        width = 4.2
        height = 3.0
        area, err := paintNeeded(width, height)
        if (err != nil) {
                fmt.Println(err)
        } else {
                PrintPaintNeeded(area)
        }
        

        width = 5.2
        height = 3.5
        area, err = paintNeeded(width, height)
        if (err != nil) {
                fmt.Println(err)
        } else {
                PrintPaintNeeded(area)
        }

        width = -15.2
        height = 3.5
        area, err = paintNeeded(width, height)
        if (err != nil) {
                fmt.Println(err)
        } else {
                PrintPaintNeeded(area)
        }

        width = 6.2
        height = -3.5
        area, err = paintNeeded(width, height)
        if (err != nil) {
                fmt.Println(err)
        } else {
                PrintPaintNeeded(area)
        }
}
