package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter an input")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Input is", input)

	fmt.Println("Enter an input")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Input is", number)
}