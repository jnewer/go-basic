package main

import "fmt"

func main() {
	s1 := []int{}
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3, 4, 5)
	fmt.Printf("s1: %v\n", s1) //s1: [1 2 3 4 5]

	s3 := []int{3, 4, 5}
	s4 := []int{1, 2}
	s4 = append(s4, s3...)
	fmt.Printf("s4: %v\n", s4) //s4: [1 2 3 4 5]
}
