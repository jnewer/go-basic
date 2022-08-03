package main

import (
	"fmt"
	"strings"
)

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func add2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	var f = add()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(20): %v\n", f(20))
	fmt.Printf("f(30): %v\n", f(30))
	fmt.Println("-----------")
	f1 := add()
	fmt.Printf("f1(40): %v\n", f1(40))
	fmt.Printf("f1(50): %v\n", f1(50))
	// f(10): 10
	// f(20): 30
	// f(30): 60
	// -----------
	// f1(40): 40
	// f1(50): 90

	var f2 = add2(10)
	fmt.Printf("f2(10): %v\n", f2(10))
	fmt.Printf("f2(20): %v\n", f2(20))
	fmt.Printf("f2(30): %v\n", f2(30))
	fmt.Println("-----------")
	f3 := add2(20)
	fmt.Printf("f3(40): %v\n", f3(40))
	fmt.Printf("f3(50): %v\n", f3(50))
	// f2(10): 20
	// f2(20): 40
	// f2(30): 70
	// -----------
	// f3(40): 60
	// f3(50): 110

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	f4, f5 := calc(10)
	fmt.Println(f4(1), f5(2)) //11 9
	fmt.Println(f4(3), f5(4)) //10 6
	fmt.Println(f4(5), f5(6)) //13 7
}
