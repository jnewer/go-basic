package main

import "fmt"

func a(n int) int {
	if n == 1 {
		return 1
	}

	return n * a(n-1)
}

func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return f(n-1) + f(n-2)
}

func main() {
	n := 5
	r := a(n)
	fmt.Printf("r: %v\n", r) //r: 120

	fab := f(5)
	fmt.Printf("fab: %v\n", fab) //fab: 5
}
