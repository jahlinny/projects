package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func containsSymbol2(symbols string, input string) bool {
	for _, char := range input {
		if strings.ContainsRune(symbols, char) {
			return true
		}
	}
	return false
}

func main() {
	symbolsToCheck2 := "@"

	for {
		scan2 := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter Your First & Last Name: ")
		scan2.Scan()
		name := (scan2.Text())

		var age int
		fmt.Print("Enter Your Age: ")
		fmt.Scan(&age)

		if age <= 18 {
			fmt.Println("You are too young")
			break
		}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter Your Email Address: ")
		scanner.Scan()
		email := (scanner.Text())

		if containsSymbol2(symbolsToCheck2, email) == false {
			fmt.Println("Please include a valid email address")
			continue
		}

		fmt.Println("Confirm Your Info: ")
		fmt.Println("Name: ", name)
		fmt.Println("Age: ", age)
		fmt.Println("Email Address: ", email)

		var conf string

		fmt.Print("TYPE YES TO CONFIRM INFO OR NO TO REDO INFO: ")
		fmt.Scan(&conf)

		if conf == "YES" {
			fmt.Println("You're all set!")
			break
		} else if conf == "NO" {
			continue
		} else {
			fmt.Println("Unspecified")
			break
		}
	}

}
