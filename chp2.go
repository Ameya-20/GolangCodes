package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)

	line := "B#ys Played well"
	fmt.Println(line)
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(line)
	fmt.Println(fixed)

}