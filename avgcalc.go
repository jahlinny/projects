package main

import (
	"fmt"
)

func main() {

	var top int
	fmt.Println("How many times do you want to append")
	fmt.Scan(&top)

	arr := []int{}

	for i := 1; i <= top; i++ {

		var app int
		fmt.Println("Type the values you want to append into the slice")
		fmt.Scan(&app)
		arr = append(arr, app)
		fmt.Println(arr)
	}
	sum := 0
	for _, num := range arr {
		sum += num

	}
	tot := float64(sum) / float64(top)
	fmt.Printf("Average: %v \n", tot)
	fmt.Printf("Total Value of integers: %v \n", sum)
}
