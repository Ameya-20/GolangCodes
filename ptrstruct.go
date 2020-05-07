package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   []string
}

func (p *Person) getFullName() string {

	//fmt.Println(p.firstName + p.lastName)
	return p.firstName + p.lastName
}

func main() {
	person := &Person{
		firstName: "Ameya",
		lastName:  "Rahurkar",
		age:       30,
		address:   []string{"Mumbai", "Paris", "London"},
	}
	//getFullName(&person)
	//variable: = &person
	fmt.Println(person.getFullName())
}
