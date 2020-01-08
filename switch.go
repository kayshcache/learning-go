// Switch statement explore

package main

import (
	"fmt"
)

func main() {
	x := 2

	switch x {
	case 1:
		fmt.Println("one")
	case 3:
		fmt.Println("three")
	case 5:
		fmt.Println("five")
	case 7:
		fmt.Println("seven")
	default:
		fmt.Printf("many")
	}

	switch {
	case x > 10:
		fmt.Println("greater than 10")
	case 10000 > 10100:
		fmt.Println("never true")
	default:
		fmt.Println("it's all we got")
	}
}
