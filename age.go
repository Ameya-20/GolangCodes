package main

import "fmt"

func main() {
	var man struct {
		name string
		age  int
	}

	man.name = "Ramlal"
	man.age = 55
	fmt.Println("Name:", man.name)
	fmt.Println("Age:", man.age)
}
