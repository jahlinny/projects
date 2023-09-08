package main

import (
	"fmt"
)

func sub(x, y float64) (z float64) {
	z = x - y
	return
}

func remainSub(x, y float64) (z float64) {
	z = x + y
	return
}

func main() {
	for {
		var total float64
		fmt.Println("Total Amount Given: ")
		fmt.Scan(&total)

		var cost float64
		fmt.Println("Total Cost of Items: ")
		fmt.Scan(&cost)

		change := sub(total, cost)

		if total < 0.00 || total == 0 {
			fmt.Println("Invalid.")
			break
		} else if change < 0.00 {
			fmt.Println("Can't return values")
			break
		} else if cost < 0.00 {
			fmt.Println("Enter correct cost")
			break
		} else if cost == 0.00 {
			fmt.Println("No cost")
			break
		} else if change == 0.00 {
			fmt.Println("No change required.")
			break
		}

		fmt.Printf("Total Change Received: %0.2f \n ", change)
		hundred := int(change / 100)
		times := float64(hundred)
		remain := (change - (times * 100))

		fifty := int(remain / 50)
		times1 := float64(fifty)
		remain1 := (remain - (times1 * 50))

		twenty := int(remain1 / 20)
		times2 := float64(twenty)
		remain2 := (remain1 - (times2 * 20))

		ten := int(remain2 / 10)
		times3 := float64(ten)
		remain3 := (remain2 - (times3 * 10))

		five := int(remain3 / 5)
		times4 := float64(five)
		remain4 := (remain3 - (times4 * 5))

		one := int(remain4 / 1)
		times5 := float64(one)
		remain5 := (remain4 - (times5 * 1))

		quarter := int(remain5 / 0.25)
		times6 := float64(quarter)
		remain6 := (remain5 - (times6 * 0.25))

		dime := int(remain6 / 0.10)
		times7 := float64(dime)
		remain7 := (remain6 - (times7 * 0.10))

		nickel := int(remain7 / 0.05)
		times8 := float64(nickel)
		remain8 := (remain7 - (times8 * 0.05))

		penny := int(remain8 / 0.01)
		times9 := float64(penny)
		remain9 := (remain8 - (times9 * 0.01))

		fmt.Printf("Hundreds: %0.2f \n", times)
		fmt.Printf("Fifties: %0.2f\n", times1)
		fmt.Printf("Twenties: %0.2f\n", times2)
		fmt.Printf("Tens: %0.2f\n", times3)
		fmt.Printf("Fives: %0.2f\n", times4)
		fmt.Printf("Ones: %0.2f\n", times5)
		fmt.Printf("Quarters: %0.2f\n", times6)
		fmt.Printf("Dimes: %0.2f\n", times7)
		fmt.Printf("Nickels: %0.2f\n", times8)
		fmt.Printf("Pennies: %0.2f\n", times9)

		fmt.Printf("Remainder ( Add this number to penny amount if it appears as 0.01 ): %0.2f", remain9)
		break

		/*if remain9 < 0.02 {
			remainSub(remain9, 0.01)
		}*/
	}
}
