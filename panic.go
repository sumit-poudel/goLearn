package main

import "fmt"

func cleanUp() {
	fmt.Println("candy cleanup")
	r := recover()
	if r == nil {
		fmt.Println(r)
	}
}

func main() {
	defer cleanUp()
	var candy uint8
	fmt.Print("Kya aap le life ma candy he? 1:No 0:Yes ")
	fmt.Scanf("%d", &candy)
	switch candy {
	case 0:
		fmt.Println("Congrulation!!")
		break
	case 1:
		panic("NO candy!")
	default:
	}
}
