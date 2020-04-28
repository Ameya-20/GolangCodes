package main

import (
	"fmt"
)

func main() {
	repeatLine("Hello World", 3)

	functionA(2,3)
	functionB(2,3)
	
}
func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}

}
func functionA(no1 int,no2 int){
	fmt.Println(no1+no2)
}

func functionB(no1 int,no2 int){
	fmt.Println(no1*no2)
}