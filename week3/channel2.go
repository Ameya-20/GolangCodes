package main

import (
	"sync"
	"fmt"
	"time"
)

var wg = sync.WaitGroup{}

func main(){
	wg.Add(2)
	ch := make(chan int,2)
	go readfromChannel(ch)
	go writetoChannel(ch)
	wg.Wait()
}

func writetoChannel(ch chan int) {
	fmt.Println("Started Writing.")
	time.Sleep(time.Second)
	ch <- 10
	ch <- 8
	ch <- 6
	fmt.Println("Data has been written to channel.")
	wg.Done()
}

func readfromChannel(ch chan int){
	fmt.Println("Started reading.")
	value := <-ch
	fmt.Println("Data in our channel is: ", value)
	wg.Done()
}