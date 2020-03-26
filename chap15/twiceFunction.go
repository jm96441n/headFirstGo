package main

import "fmt"

func SayHi() {
        fmt.Println("Say Hi")
}

func SayBye() {
        fmt.Println("Bye Felicia")
}

func Divide(a int, b int) float64 {
        return float64(a) / float64(b)
}

func Multiply(a int, b int) float64 {
        return float64(a) * float64(b)
}

func Twice(theFunc func()) {
        for i := 0; i < 2; i++ {
                theFunc()
        }
}

func doMath(divider func(int, int) float64) {
        res := divider(10, 5)
        fmt.Printf("%.1f\n", res)
}

func main() {
        var greeterFunction func()
        var dividerFunction func(int, int) float64
        greeterFunction = SayHi
        dividerFunction = Divide
        Twice(greeterFunction)
        Twice(SayBye)
        doMath(dividerFunction)
        doMath(Multiply)
}
