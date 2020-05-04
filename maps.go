package main

import "fmt"

func main() {
	body := make(map[string]float64)

	body["hands"] = 2
	body["handfingers"] = 10

	my, ok := body["handfingers"]

	fmt.Println(my, ok)

	cloth := map[string]float64{"pants": 70.6, "shirt": 40.9}
	{
		fmt.Println(jewel, cloth)
	}

}
