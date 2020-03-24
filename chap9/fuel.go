package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(0.264 * l)
}

func (l Liters) ToMilliliters() Milliliters {
	return Milliliters(1000 * l)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (m Milliliters) ToLiters() Liters {
	return Liters(0.001 * m)
}

func (g Gallons) ToLiters() Liters {
	return Liters(3.785 * g)
}

func main() {
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)
	fmt.Println(carFuel, busFuel)
	carFuel = Gallons(Liters(40.0) * 0.264)
	busFuel = Liters(Gallons(63.0) * 3.785)
	fmt.Printf("Gallons: %0.1f, Liters: %0.1f\n", carFuel, busFuel)

	carFuel2 := Gallons(1.2)
	busFuel2 := Liters(4.5)
	carFuel2 += Liters(40.0).ToGallons()
	busFuel2 += Gallons(30.0).ToLiters()
	fmt.Printf("CarFuel: %0.1f\n", carFuel2)
	fmt.Printf("BusFuel: %0.1f\n", busFuel2)

	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}
