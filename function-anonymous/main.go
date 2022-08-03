package main

import "fmt"

func main() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	i := max(1, 2)
	fmt.Printf("i: %v\n", i) //i: 2

	func(a int, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max: %v\n", max) //max: 2
	}(1, 2)
}
