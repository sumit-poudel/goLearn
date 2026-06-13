package main

import "fmt"

type Number interface {
	int | float64 | uint
}

func add[T Number](a T, b T) T {
	return a + b
}

func mapIterator[C comparable, V any](mp map[C]V) []V {
	var values []V
	for _, value := range mp {
		values = append(values, value)
	}
	return values
}

type values interface { // <-- cant to values[T any] it will loose ability to have diffrent type slice ie. []any {int,string}
	getValue() any
}

type genStruct[T any] struct {
	value T
}

func (r genStruct[T]) getValue() any { //<-- inmplements interface
	return r.value
}

func main() {
	fmt.Println("The sum is ", add(4, 5))
	fmt.Println("The sum is ", add(3.4, 5.7))
	fmt.Println("The sum is ", add(uint(3), uint(5)))
	fmt.Println("The values are", mapIterator(map[int]string{
		1: "sumit", 2: "poudel", 3: "ram lal",
	}))
	person := []genStruct[any]{{value: "sumit poudel"}, {value: 22}}
	for _, value := range person {
		fmt.Println(value.value)
	}

	nextPerson := []values{genStruct[string]{value: "samir poudel"}, genStruct[int]{value: 4}}
	for _, nextValue := range nextPerson {
		fmt.Println(nextValue.getValue())
	}
}
