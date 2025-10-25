package main

import (
	"fmt"
	"reflect"
)

func varType() {
	var a string
	var b int
	var c bool

	fmt.Println(a, "Type of a:", reflect.TypeOf(a))
	fmt.Println(b, "Type of b:", reflect.TypeOf(b))
	fmt.Println(c, "Type of c:", reflect.TypeOf(c))
}
