package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Errror Occured")
	}
}

func main() {
	data := []byte("Hello from Ameya Rahurkar")
	err := ioutil.WriteFile("myfile.txt", data, 0644)
	checkError(err)
	fmt.Println("\n\n\nDONE WRITING")

	f, er := os.Create("myfile2.txt")
	checkError(er)
	defer f.Close()

	bytesWritten,
	e := f.WriteString("Hello Namastey!!!")
	checkError(e)

	fmt.Println("Bytes Written are: ", bytesWritten)
	fmt.Println("\n\n\nDONE WRITING")

}
