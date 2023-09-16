package main

import (
	"fmt"
)

func sub(x, y float64) (z float64) {
	z = x - y
	return
}

func main() {

	fmt.Println("Slope Calculator - ( Y2 - Y1 ) / ( X2 - X1 )")

	var y2 float64
	fmt.Println("Type Y2 value: ")
	fmt.Scan(&y2)

	var y1 float64
	fmt.Println("Type Y1 Value: ")
	fmt.Scan(&y1)

	var x2 float64
	fmt.Println("Type X2 Value: ")
	fmt.Scan(&x2)

	var x1 float64
	fmt.Println("Type X1 Value: ")
	fmt.Scan(&x1)

	tot1 := sub(y2, y1)
	tot2 := sub(x2, x1)

	total1 := tot1 / tot2

	fmt.Printf("Slope in Fraction Form: %0.2f/%0.2f \n", tot1, tot2)
	fmt.Printf("Slope in Numerical Form: %0.2f", total1)

}
