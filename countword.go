package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Paste the text: ")
	scanner.Scan()
	input := (scanner.Text())
	count := strings.Count(input, " ")

	if count == 0 {
		fmt.Println("There is", count+1, "word in this text")
	} else if count >= 1 {
		fmt.Println("There are ", count+1, "words in this text")
	}

}
