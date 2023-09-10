package main

import (
	"fmt"
)

func add(x, y float64) (z float64) {
	z = x + y
	return
}

func main() {
	var num float64
	fmt.Print("Type in your starting number: ")
	fmt.Scan(&num)
	count := 1
	for {
		var conf string
		fmt.Println("Type YES to continue adding numbers to your average or type DONE to find your average based off of your number inputs")
		fmt.Scan(&conf)

		if conf == "YES" {
			var num1 float64
			fmt.Print("Type in the number you want to add: ")
			fmt.Scan(&num1)
			num = add(num, num1)
			fmt.Println("Total: ", num)
			count++
			continue
		} else if conf != "DONE" {
			fmt.Println("Please enter DONE to calculate the average or YES to keep adding values")
			continue
		} else if conf == "DONE" {
			if count == 1 {
				fmt.Println("No numbers added")
				break
			} else if count > 1 {

				fmt.Println("Average: ", num/float64(count))
				break
			}
		}

	}
}
