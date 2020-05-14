package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("CLose Resourse")
		details := recover()
		fmt.Println("Details : ", details)
	}()
	fmt.Println("Open Resource")
	panic("I dont know what to do")
	fmt.Println("End of main")
}
