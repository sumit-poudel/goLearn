package main

import "fmt"

func changeByValue(a int,b int){
	a = b
}

func changeByRefrence(a *int,b int){
	*a=b
}

func main(){
	value1:=4
	value2:= 5
	fmt.Println("The value before change is a: ",value1,"b: ",value2)
	changeByValue(value1,value2)
	fmt.Println("The value after change is a: ",value1,"b: ",value2)
	changeByRefrence(&value1,value2)
	fmt.Println("The value after change is a: ",value1,"b: ",value2)
}
