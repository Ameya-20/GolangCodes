package main

import (
	"fmt"
)

type whistle string

func (w whistle) makesound() {
	fmt.Println("Tweet!")
}

type horn string

func (h horn) makesound() {
	fmt.Println("Honk!")
	fmt.Println(h)
}

type noiseMaker interface {
	makesound()
}

func main() {
	var toy noiseMaker
	toy = whistle("canary")
	toy.makesound()
	toy = horn("blaster")
	toy.makesound()
}
