package main

import "fmt"

func main(){
	var num=100
	f1(num)
	defer f2()
	num=50
	fmt.Println("End of main")
	fmt.Println(num)

}

func f1(num int){
	defer fmt.Println("Defer in f1")
	fmt.Println("End of f1")
	fmt.Println(num)
}

func f2(){
	fmt.Println("End of f2")
}