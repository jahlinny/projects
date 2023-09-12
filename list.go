package main

import (
	"fmt"
)

func main() {

	var list int
	fmt.Print("How many numbers do you want on your list?: ")
	fmt.Scan(&list)

	for i := 0; i <= list; i++ {
		fmt.Printf("%v. \n", i)
	}
}
