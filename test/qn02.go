package main

import (
	"errors"
	"fmt"
)

// YOUR CODE HERE:
// Declare a "divide" function such that the call in the
// "main" function will compile and return 2.8.
// "divide" should accept two float64 values as parameters,
// and return a single float64 value that represents the
// first parameter divided by the second.
// EXTRA CREDIT:
// Have "divide" return TWO values, a float64 and an error.
// If the second parameter is 0, return an error value
// with the message "can't divide by 0". Otherwise, return
// nil for the error value. You can use the fmt.Errorf
// function to generate an error value. You'll also need
// to update the code in "main" to handle the error value.

func divide(no1, no2 float64) (float64, error) {
	if no2 == 0 {

		return 0, errors.New("Cannot divide by 0!! ")
	} else {
		return no1 / no2, nil
	}
}

func main() {
	quotient, error := divide(5.6, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%0.2f\n", quotient) // => 2.80
	}

}
