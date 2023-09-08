package main

import (
	"fmt"
	"math"
)

func div(x, y float64) (z float64) {
	z = x / y
	return
}

func mult(x, y float64) (z float64) {
	z = x * y
	return
}
func add(x, y float64) (z float64) {
	z = x + y
	return
}
func sub(x, y float64) (z float64) {
	z = x - y
	return
}
func exp(x, y float64) (z float64) {
	z = math.Pow(x, y)
	return
}

func main() {
	var num float64
	fmt.Println("Type the number you want to perform your operations on")
	fmt.Scan(&num)
	for {
		var sign string
		var calc float64

		fmt.Println("Type the operand symbol (+, -, *, /, ^) and the number you want to operate the input on seperate lines ( ex: / (enter) 9 (enter) = /9 ): ")
		fmt.Scan(&sign, &calc)

		switch sign {
		case "/":
			num = div(num, calc)
		case "*":
			num = mult(num, calc)

		case "+":
			num = add(num, calc)
		case "-":
			num = sub(num, calc)
		case "^":
			num = exp(num, calc)
		default:
			fmt.Println("Sign unavailable, please try again with /, *, +, -, ^")
			continue
		}
		fmt.Println("Current Total: ", num)
		var confirm string
		fmt.Printf("To continue your operation, type YES, to break out of the program, type NO: ")
		fmt.Scan(&confirm)

		if confirm == "NO" {
			fmt.Println("Goodbye")
			break
		}

		/*if strings.Contains(sign, "/") {
			fmt.Println("Answer: ", div(num, calc))
		} */

	}
}
