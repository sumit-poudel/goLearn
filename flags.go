package main

import (
	"flag"
	"fmt"
)

func main() {
	name :=flag.String("name","user","Your name flag")
	flag.Parse()
	fmt.Printf("Hello %s, My name is go.",*name)
}
