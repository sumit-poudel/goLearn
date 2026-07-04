package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	connectionWithTimeout()
}

func openConnection(done chan bool) {
	fmt.Println("Attempting Connection....")

	if rand.Intn(100) > 50 {
		fmt.Println("OPPS! Hanging Connection!")
		time.Sleep(10000 * time.Hour)
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Connection Established!")
	}
	done <- true
}

func connectionWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	done := make(chan bool)
	go openConnection(done)

	select {
	case <-done: // this return a bool form chan
		fmt.Println("connection Successful")
	case <-ctx.Done():     // this returns a chan {}
		fmt.Println("connection Timeout!")
	}
}
