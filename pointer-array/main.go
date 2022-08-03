package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{1, 2, 3}
	var i int
	var ptr [MAX]*int
	fmt.Println(ptr)
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	// [<nil> <nil> <nil>]
	// a[0] = 1
	// a[1] = 2
	// a[2] = 3
}
