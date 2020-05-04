package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")

	for _, num := range nums {
		fmt.Println(num)
		println(num)
	}

}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
