package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the guessing game!")

	var guess int
	fmt.Print("Choose a number 1-10: ")
	input := readInput2()
	guess, err := strconv.Atoi(input)
	if err != nil || guess < 1 || guess > 10 {
		fmt.Println("Invalid input")
		return
	}

	bguess := rand.Intn(10)

	for {
		var next string
		fmt.Print("Do you want to see what the bot chose? ( type YES to continue or NO to exit): ")
		fmt.Scan(&next)

		if next == "YES" {
			fmt.Println("The bot chose: ", bguess)
			if bguess == guess {
				fmt.Println("You won!")
				break
			} else if bguess != guess {
				fmt.Println("You lost!")
				break
			}
		} else if next == "NO" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("YES or NO only")
			break
		}

	}
}

func readInput2() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	return input
}
