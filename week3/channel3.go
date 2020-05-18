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
	ch2 := make(chan struct{})
	go readfromChannel(ch)
	go writetoChannel(ch)
	wg.Wait()
	close(ch)
	go func()  {
		ch2 <- struct{}{}
	}() 

	for i  := range ch{
		fmt.Println("Inside range ",i)
	}
	for i:= 0 ; i<2 ; i++{
		select{
		case chan1,ok := <- ch:
			fmt.Println("Channel 1 has data", chan1,ok)
		case chan2,ok := <- ch2:
			fmt.Println("Channel 2 has data", chan2,ok)
				
		}
	}
}

func writetoChannel(ch chan<- int) {
	fmt.Println("Started Writing.")
	time.Sleep(time.Second)
	ch <- 10
	ch <- 8
	ch <- 6
	fmt.Println("Data has been written to channel.")
	wg.Done()
}

func readfromChannel(ch <-chan int){
	fmt.Println("Started reading.")
	value,ok := <-ch
	value2 := <-ch
	fmt.Println("Data in our channel is: ", value,value2,ok)
	wg.Done()
}