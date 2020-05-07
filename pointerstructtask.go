package main

import "fmt"

type Persondata struct {
	firstName string
	lastName  string
	age       int
	address   []string
}

func printdata(p *Persondata) {
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
	fmt.Println(p.age)

	for i := 0; i < len(p.address); i++ {
		fmt.Println(p.address[i])
	}
}

func main() {
	person := Persondata{
		firstName: "Ameya",
		lastName:  "Rahurkar",
		age:       21,
		address:   []string{"Mumbai", "Paris", "London"},
	}
	printdata(&person)
}
