package main

import "fmt"


func main() {
	
	var option int
	
	fmt.Println("Choose\n1.ADD\n2.Subtract\n3.Multiply\n4.Divide\n\nOption- ")
	fmt.Scanf("%d", &option)
	
	if option == 1 {
		fmt.Println("Answer = ",mathOperation(add))
	}
	
	if option == 2 {
		fmt.Println("Answer = ",mathOperation(subtract))
	}
	if option == 3 {
		fmt.Println("Answer = ",mathOperation(multiply))
	}

	if option == 4 {
		fmt.Println("Answer = ",mathOperation(div))
	}
	
}

func mathOperation(funcName func(num1, num2 float64)float64)float64 {
	 
	return funcName(100,10)
	 
}

func add(no1 float64, no2 float64) float64 {

	return (no1 + no2)
}

func subtract(no1 float64, no2 float64) float64 {

	return (no1 - no2)
}

func div(no1 float64, no2 float64) float64 {

	return (no1 / no2)
}

func multiply(no1 float64, no2 float64) float64 {

	return (no1 * no2)
}
