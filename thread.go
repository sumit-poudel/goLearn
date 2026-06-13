package main

import (
	"fmt"
	"sync"
	"time"
)

func run() {
	time.Sleep(1 * time.Second)
	fmt.Println("run")
}
func run2() {
	time.Sleep(2 * time.Second)
	fmt.Println("run")
}

func addition(a int, b int, ch chan int) {
	ch <- a + b
}
func wgAddition(a int,b int,wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(a+b)
}

func main() {
	fmt.Println("Threads: ")
	ch := make(chan int) //<- channel to receive data from routine
	bch :=make(chan int,1) // <- buffered channel saves 1 data type int
	go addition(3, 4, ch) //<- regular gochannel so have to wait for its output
	addition(68,1,bch) //<- buffered channel saves output in buffer
	x := <-ch
	fmt.Println(x)
	fmt.Println("Buffered channel")
	y:=<-bch
	fmt.Println(y)
	fmt.Println("wait group")
	wg:= sync.WaitGroup{}
	wg.Add(2)
	wgAddition(2,3,&wg)
	wgAddition(5,6,&wg)
	/* creates go routine
	go run()
	go run2()
	time.Sleep(2 * time.Second)
	*/
	wg.Wait()
}
