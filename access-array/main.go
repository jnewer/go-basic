package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200

	fmt.Printf("a[0]: %v\n", a[0]) //a[0]: 100
	fmt.Printf("a[1]: %v\n", a[1]) //a[1]: 200

	a[0] = 1
	a[1] = 2
	fmt.Println("-----------")

	fmt.Printf("a[0]: %v\n", a[0]) //a[0]: 1
	fmt.Printf("a[1]: %v\n", a[1]) //a[1]: 2

	aa := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(aa); i++ {
		fmt.Printf("aa[%d]: %v\n", i, aa[i])
	}
	// aa[0]: 1
	// aa[1]: 2
	// aa[2]: 3
	// aa[3]: 4
	// aa[4]: 5
	// aa[5]: 6

	for i, v := range aa {
		fmt.Printf("aa[%d]: %v\n", i, v)
	}
	// aa[0]: 1
	// aa[1]: 2
	// aa[2]: 3
	// aa[3]: 4
	// aa[4]: 5
	// aa[5]: 6

	
}
