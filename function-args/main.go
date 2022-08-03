package main

import "fmt"

/* func f1(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
} */

/* func f1(a int) {
	a = 200
	fmt.Printf("a1: %v\n", a)
} */

/* func f1(a []int) {
	a[0] = 100
} */

func f1(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func f2(name string, age int, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	// r := f1(1, 2)
	// fmt.Printf("r: %v\n", r)

	// a := 100
	// f1(a)
	// fmt.Printf("a: %v\n", a)

	// a := []int{1, 2}
	// f1(a)

	f1(1, 2, 3)
	fmt.Println("------------")
	f1(1, 2, 3, 4, 5, 6)
	fmt.Println("------------")
	f2("tom", 20, 1, 2, 3)
}
