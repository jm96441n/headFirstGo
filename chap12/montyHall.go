package main

import (
        "fmt"
        "math/rand"
        "time"
)
func awardPrize() {
        doorNumber := rand.Intn(3) + 1
        if doorNumber == 1 {
                fmt.Println("You win a cruise!")
        } else if doorNumber == 2 {
                fmt.Println("You win a car!")
        } else if doorNumber == 3 {
                fmt.Println("You win a goat!")
        } else {
                panic("Haven't you people ever heard of closing the goddamn door")
        }
}

func main() {
        rand.Seed(time.Now().Unix())
        awardPrize()
}
