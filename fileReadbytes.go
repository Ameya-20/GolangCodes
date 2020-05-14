package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(content)
	result := string(content)
	fmt.Println(result)

	f, er := os.Open("myfile.txt")
	defer f.Close()

	checkError(er)
	bucket := make([]byte, 100)
	bytesRead, e := f.Read(bucket)
	checkError(e)
	fmt.Println("\n\n\n\nContent of file(limited) Read: ", string(bucket[:bytesRead]))
	fmt.Println("Bytes of file read: ", bytesRead)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error Occured!!", err)

	}
}
