package main

import (
	"fmt"
	"sync"
	"time"
)

var m int = 100
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	lock.Lock()
	m += 1
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	defer wg.Done()
	lock.Lock()
	m -= 1
	time.Sleep(time.Millisecond * 2)
	lock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		wg.Add(1)
		go sub()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("m: %v\n", m) //m: 100
}
