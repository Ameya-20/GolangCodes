package main

import "fmt"

func main() {
	loopfunct("list Starts now", 5)
}

func loopfunct(str string, counter int) {
	for i := 1; i <= counter; i++ {
		fmt.Println(str)
	}
}
