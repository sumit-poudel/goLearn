package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func readConfig(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg map[string]any
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
func main() {

	defer func() {
		errr := recover()
		if errr != nil {
			log.Fatal(errr)
		}
	}()
	filePath := flag.String("path", "./data.json", "your config file")
	flag.Parse()
	fmt.Println("path to file is ", *filePath)

	data, err := readConfig(*filePath)
	if _, ok := errors.AsType[*fs.PathError](err); ok {
		panic("path error")
	}
	if _, ok := errors.AsType[*json.SyntaxError](err); ok {
		panic("syntax error")
	}
	fmt.Printf("%+v", data)
}
