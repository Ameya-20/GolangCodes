package main

import (
	"errors"
	"fmt"
)

func divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return numerator / denominator, nil
}

func main() {
	quotient, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(quotient)
	}
}
