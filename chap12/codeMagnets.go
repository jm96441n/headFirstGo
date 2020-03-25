package main

import (
        "fmt"
        "log"
)

func find(item string, slice []string) bool {
        found := false
        for _, sliceItem := range slice {
               if sliceItem == item {
                        found = true
                        break
               }
        }
        return found
}

type Regfrigerator []string

func (r Regfrigerator) Open() {
        fmt.Println("Opening regfrigerator")
}

func (r Regfrigerator) Close() {
        fmt.Println("Closing regfrigerator")
}

func (r Regfrigerator) FindFood(food string) error {
        r.Open()
        defer r.Close()
        if find(food, r) {
                fmt.Println("Found", food)
        } else {
                return fmt.Errorf("%s is not in regfrigerator\n", food)
        }
        return nil
}

func main() {
        fridge := Regfrigerator{"Milk", "Pizza", "Salsa"}
        for _, food := range []string{"Milk", "Bananas"} {
                err := fridge.FindFood(food)
                if err != nil {
                        log.Fatal(err)
                }
        }
}
