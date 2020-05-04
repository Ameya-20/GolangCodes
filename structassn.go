package main

import "fmt"

type person struct {
	name   string
	age    float64
	active bool
}

func main() {
	var person1 person
	var person2 person

	person1.name = "Ameya"
	person2.name = "Raghav"

	person1.rate = 83.5
	person2.rate = 22.5

	person1.active = false
	person2.active = true

	fmt.Println(person1, person2)
}
