// Package main by sumit
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	age := "21"
	ageInt, _ := strconv.Atoi(age)
	fmt.Println("The age is ", ageInt, " the type is ", reflect.TypeOf(ageInt))
	name := "Su🥳mit Poudel"
	for idx := 0; idx < len(name); idx++ {
		fmt.Printf("%c", name[idx])
	}
	fmt.Println()
	for _, char := range name {
		fmt.Printf("%c", char)
	}
}
