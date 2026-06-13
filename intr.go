package main

import "fmt"

type animal interface {
	breath() string
}
type duck interface {
	animal
	fly() string
}

type raja struct {
	name string
}

func (d raja) fly() string {
	return d.name + " fly in the sky"
}

func (d raja) breath() string {
	return d.name + " breaths air"
}

func main() {
	var ducky raja = raja{name: "donald duck"}
	fmt.Println(ducky.name, "is a duck", ducky.breath(), ducky.fly())
}
