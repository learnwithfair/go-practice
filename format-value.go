package main

import (
	"fmt"
)

func valueFormat() {
	i, j := "Hello", 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
