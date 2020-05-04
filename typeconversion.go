package main

import "fmt"

func main() {
	var a = 23
	var b = 100.54
	ans := float64(a) * b
	fmt.Println("Int a = 23 multiplied by Float b = 100.54 is", ans)
	var c = 15.0
	var d = 17
	answer := int(c) * d
	fmt.Println("Float c = 25.0 multiplied by int d = 10 is", answer)
}
