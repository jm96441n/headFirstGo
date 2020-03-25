package main

import "fmt"

type Car string

func (c Car) Accelerate() {
        fmt.Println("Speeding up!")
}

func (c Car) Brake() {
        fmt.Println("Braking!")
}

func (c Car) Steer(direction string) {
        fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
        fmt.Println("Speeding up!")
}

func (t Truck) Brake() {
        fmt.Println("Braking!")
}

func (t Truck) Steer(direction string) {
        fmt.Println("Turning", direction)
}

func (t Truck) LoadCargo(cargo string) {
        fmt.Println("Loading", cargo)
}

type Vehicle interface {
        Accelerate()
        Brake()
        Steer(string)
}

func TryVehicle (vehicle Vehicle) {
        vehicle.Accelerate()
        vehicle.Steer("Left")
        vehicle.Steer("Right")
        vehicle.Brake()
        truck, ok := vehicle.(Truck)
        if ok {
                truck.LoadCargo("Test Cargo")
        }
}

func main() {
        TryVehicle(Truck("Ford F180"))
}
