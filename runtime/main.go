package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func show2() {
	for i := 0; i < 10; i++ {
		if i >= 5 {
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", i)
	}
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	// go show("java")

	// for i := 0; i < 2; i++ {
	// 	runtime.Gosched()
	// 	fmt.Println("golang")
	// }

	// go show2()
	// time.Sleep(time.Second)

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
