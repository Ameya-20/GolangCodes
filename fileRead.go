package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
	result := string(content)
	fmt.Println(result)

}
