// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {
	x := 11.0
	y := 20.0

	if frac := x / y; frac > 0.5 {
		fmt.Println("x is more than half of y")
	}

}
