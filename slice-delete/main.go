package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1) //s1: [1 2 4 5]
}
