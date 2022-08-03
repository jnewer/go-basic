package main

import "fmt"

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func compare(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}

	return max
}
func main() {
	s := sum(1, 2)
	fmt.Printf("s: %v\n", s) //s: 3

	max := compare(1, 2)
	fmt.Printf("max: %v\n", max) //max: 2
}
