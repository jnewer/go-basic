package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[0:3]
	fmt.Printf("s2: %v\n", s2) //s2: [1 2 3]
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3) //s3: [4 5 6]
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4) //s4: [3 4 5]
	s5 := s1[:]
	fmt.Printf("s5: %v\n", s5) //s5: [1 2 3 4 5 6]

	s := []int{1, 2, 3}
	fmt.Printf("s: %v\n", s) //s: [1 2 3]
	arr := [...]int{1, 2, 3}
	sarr := arr[:]
	fmt.Printf("sarr: %v\n", sarr) //sarr: [1 2 3]

	var nils []int
	fmt.Println(nils == nil)                              //true
	fmt.Printf("len: %d cap: %d\n", len(nils), cap(nils)) //len: 0 cap: 0

	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%d]: %d\n", i, s1[i])
	}

	for i, v := range s1 {
		fmt.Printf("s1[%d]: %d\n", i, v)
	}
}
