package main

import "fmt"

type Person struct {
	name  string
	roll  int
	plays []Sports
}

type Sports struct {
	name   string
	number int
}

func main() {
	var sumit Person = Person{"sumit", 24, []Sports{{"cricket", 20}, {"football", 20}}}
	fmt.Println(sumit.name + sumit.plays[1].name)
}
