package main

import "fmt"

type myInterface interface {
	methodWithoutParmameters()
	methodWithParmameters(float64)
	methodWithReturn() string
}

type myType int

func (m myType) methodWithoutParmameters() {
	fmt.Println("methodWithoutParmameters called")
}

func (m myType) methodWithParmameters(float64) {
	fmt.Println("methodWithParmameters called")
}

func (m myType) methodWithReturn() string {
	return "methodWithReturn called"
}

func main() {
	var value myInterface
	value = myType(5)
	value.methodWithoutParmameters()
	value.methodWithParmameters(127.3)
	fmt.Println(value.methodWithReturn())
}
