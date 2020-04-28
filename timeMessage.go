package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()

	if currentTime.Hour() < 12 {
		fmt.Println("Good Morning!!")
	} else if currentTime.Hour() < 16 {
		fmt.Println("Good AfterNoon!!")
	} else if currentTime.Hour() < 19 {
		fmt.Println("Good Evening!!")
	} else if currentTime.Hour() >= 19 {
		fmt.Println("Good Day!!")
	} else {
		fmt.Println("Good Night!!")
	}

}
