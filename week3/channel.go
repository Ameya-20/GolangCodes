package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go writeToChannel(ch)
	value := <-ch
	fmt.Println("Data in our channel is: ", value)
	go writeToChannel2(ch2)
	value2 := <-ch2
	fmt.Println("Data in our channel is: ", value2)

}

func writeToChannel(ch chan int) {
	fmt.Println("Inside go routine.")
	time.Sleep(time.Second)
	ch <- 10
	fmt.Println("Data has been written to channel.")
}

func writeToChannel2(ch2 chan int) {
	fmt.Println("Inside go routine.")
	ch2 <- 20
	time.Sleep(time.Second)
	fmt.Println("Data has been written to channel.")
}
