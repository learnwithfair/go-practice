package main

import (
	"fmt"
)

func conditionsDemo() {
	a := 10
	b := 5
	c := 10

	// Largest number calculation
	if a >= b && a >= c {
		fmt.Println("a is the largest number")
	} else if b >= a && b >= c {
		fmt.Println("b is the largest number")
	} else {
		fmt.Println("c is the largest number")
	}
}
