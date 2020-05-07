package main

import "fmt"

type Car string

func (c Car) accelerate() {
	fmt.Println("Speeding up")
}

func (c Car) brake() {
	fmt.Println("Stopping")
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type truck string

func (t truck) accelerate() {
	fmt.Println("Speeding up")
}

func (t truck) brake() {
	fmt.Println("Stopping")
}

func (t truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t truck) loadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	accelerate()
	brake()
	Steer(direction string)
}

func main() {
	var vehicle Vehicle = Car("Toyota")
	vehicle.accelerate()
	vehicle.Steer("left")

	vehicle = truck("fnord")
	vehicle.brake()
	vehicle.Steer("right")
}
