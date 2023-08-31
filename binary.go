package main

import (
	"fmt"
	"strconv"
)

func main() {
	var binary string
	fmt.Println("Enter Binary Value: ")
	fmt.Scan(&binary)

	decimalValue, err := strconv.ParseInt(binary, 2, 32)
	if err == nil {
		fmt.Println("Parsed binary value:", decimalValue)
	} else if err != nil {
		fmt.Println("Unspecified Value, please insert a different value")
	}
}
