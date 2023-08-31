package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func containsSymbol(symbols string, input string) bool {
	for _, char := range input {
		if strings.ContainsRune(symbols, char) {
			return true
		}
	}
	return false
}

func main() {
	symbolsToCheck := "!@#$%^&*()"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Password: ")
	scanner.Scan()
	input := (scanner.Text())

	if len(input) <= 10 && containsSymbol(symbolsToCheck, input) {
		fmt.Println("Your password is mediocre")
	} else if len(input) > 12 && len(input) < 20 && containsSymbol(symbolsToCheck, input) {
		fmt.Println("Your password is strong")
	} else if len(input) >= 20 && containsSymbol(symbolsToCheck, input) {
		fmt.Println("Your password is really strong")
	} else if len(input) >= 20 {
		fmt.Println("Your password is mediocre, try adding some symbols ")
	} else if len(input) <= 12 {
		fmt.Println("Your password is weak")
	} else {
		fmt.Println("Unable to process")
	}
}
