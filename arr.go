package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [5]int{1, 3, 2, 4, 5}
	fmt.Println("array 1: ", arr1)
	fmt.Println("array 2: ", arr2)
	sl := arr2[2:3]
	sl = sl[:3]
	jaggedArr := [...][]int{{1, 2, 4}, {2, 3}, {5, 6, 7}}
	fmt.Println("The jagged arr is: ", jaggedArr)
	fmt.Println("The slice is: ", sl, len(sl), cap(sl))
	// map
	map1 := make(map[int]string)
	var map2 map[string][]int = map[string][]int{"a": {1, 2, 4}, "b": {14, 5, 6, 7}, "c": {2, 4, 5, 6, 3, 6, 3}}
	fmt.Println("Map: ", map1)

	fmt.Println("Map: ", map2)
}
