package main

import (
	"fmt"
)

func arrayDemo() {
	var arr1 = [...]int{2, 1, 5, 3, 4}
	arr2 := [...]string{"two", "one", "five", "three", "four"}

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2[len(arr2)-1])

}
