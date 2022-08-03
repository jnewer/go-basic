package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
	go func() {
		chanInt <- 100
		chanStr <- "hello"
		close(chanInt)
		close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
		default:
			fmt.Println("default...")
		}

		time.Sleep(time.Second)
	}
}
