package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1) //s1: [100 2 3]
	fmt.Printf("s2: %v\n", s2) //s2: [100 2 3]
	fmt.Println("----------")
	s3 := make([]int, 3)
	copy(s3, s1)
	s1[0] = 1
	fmt.Printf("s1: %v\n", s1) //s1: [1 2 3]
	fmt.Printf("s3: %v\n", s3) //s3: [100 2 3]

}
