package main

import "fmt"

type Car string

func (c Car) Accelerate() {
        fmt.Println("Speeding up!")
}

func (c Car) Brake() {
        fmt.Println("Braking!")
}

func (c Car) Turn(direction string) {
        fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
        fmt.Println("Speeding up!")
}

func (t Truck) Brake() {
        fmt.Println("Braking!")
}

func (t Truck) Turn(direction string) {
        fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
        fmt.Println("Loading", cargo)
}

type Vehicle interface {
        Accelerate()
        Brake()
        Turn(string)
}

func main() {
        var vehicle Vehicle = Car("Toyota Yarvic")
        vehicle.Accelerate()
        vehicle.Brake()

        vehicle = Truck("Ford F180")
        vehicle.Accelerate()
        vehicle.Brake()
}
