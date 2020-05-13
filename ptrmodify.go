package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func modifyPointer(p *Person) {
	p.firstName = "Ameya"
	p.lastName = "Rahurkar"
}

func modifyValue(sub Person) {
	sub.firstName = "Nivedita"
	sub.lastName = "R"
	fmt.Println(sub)
}

func main() {
	person := Person{
		firstName: "Prachi",
		lastName:  "Rahurkar",
	}
	fmt.Println(person) //original value printed
	modifyValue(person)
	fmt.Println(person) //again original as variable sub is modified, not person
	modifyPointer(&person)
	fmt.Println(person) //new values printed
}
