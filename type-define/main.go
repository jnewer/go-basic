package main

import "fmt"

func main() {
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i: %v i: %T\n", i, i) //i: 100 i: main.MyInt

	type MyInt2 = int
	var i2 MyInt2
	i2 = 100
	fmt.Printf("i: %v i: %T\n", i2, i2) //i: 100 i: int
}
