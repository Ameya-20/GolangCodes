package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	fmt.Println(10 > 7)
	fmt.Println(10 > 19)

	fmt.Println(3.5 / 6)

	fmt.Println(strings.Title("shi shuan"))

	fmt.Println(reflect.TypeOf(4))
	fmt.Println(reflect.TypeOf(4.9))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("jesdfxcgvh"))

	var quantity int
	quantity = 10
	//can be done this way also
	var length, breadth = 31.2, 26.4
	var custname string = "Ameya"
	fmt.Println(custname, "needs ", quantity, " boxes of tiles with for his place of ", length*breadth, " sq.ft.")
	//when variable name is initialized and no value is assigned, int,float var has
	// 0 value, string has empty string, bool has false.

	//Another way of performing same above stuff, this does initializing and assigning
	// at the same time, cant be done on predeclared variables.
	quantit := 40
	lengt, widt := 30, 60
	cusname := "Ameya"
	fmt.Println(cusname, " needs ", quantit, " boxes of tiles with for his place of ", lengt*widt, " sq.ft.")

	inte := 2
	flot := 4.0
	//we need type conversions if we want to apply some functions in which they
	// interact directly.
	ans := int(flot) * inte
	fmt.Println(ans)

}
