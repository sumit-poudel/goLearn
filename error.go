package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(2, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The divided is ", result)
}
