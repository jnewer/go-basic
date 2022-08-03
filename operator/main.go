package main

import "fmt"

func main() {
	a := 100
	b := 10
	fmt.Printf("(a + b): %v\n", (a + b))  //(a + b): 110
	fmt.Printf("(a - b): %v\n", (a - b))  //(a - b): 90
	fmt.Printf("(a * b): %v\n", (a * b))  //(a * b): 1000
	fmt.Printf("(a / b): %v\n", (a / b))  //(a / b): 10
	fmt.Printf("(a %% b): %v\n", (a % b)) //(a % b): 0

	a++
	fmt.Printf("a: %v\n", a) //a: 101
	b--
	fmt.Printf("b: %v\n", b) //b: 9

	c := 1
	d := 2
	fmt.Printf("(c > d): %v\n", (c > d))   //(c > d): false
	fmt.Printf("(c < d): %v\n", (c < d))   //(c < d): true
	fmt.Printf("(c >= d): %v\n", (c >= d)) //(c >= d): false
	fmt.Printf("(c <= d): %v\n", (c <= d)) //(c <= d): true
	fmt.Printf("(c == d): %v\n", (c == d)) //(c == d): false
	fmt.Printf("(c != d): %v\n", (c != d)) //(c != d): true

	e := true
	f := false
	fmt.Printf("(e && f): %v\n", (e && f)) //(e && f): false
	fmt.Printf("(e || f): %v\n", (e || f)) //(e || f): true
	fmt.Printf("(!e): %v\n", (!e))         //(!e): false
	fmt.Printf("(!f): %v\n", (!f))         //(!f): true

	g := 4
	fmt.Printf("g: %b\n", g) //g: 100
	h := 8
	fmt.Printf("h: %b\n", h) //h: 1000

	fmt.Printf("(g & h): %v, %b\n", (g & h), (g & h)) //(g & h): 0, 0
	fmt.Printf("(g | h): %v, %b\n", (g | h), (g | h)) //(g | h): 12, 1100
	fmt.Printf("(g ^ h): %v, %b\n", (g ^ h), (g ^ h)) //(g ^ h): 12, 1100
	fmt.Printf("(g << 2): %v\n", (g << 2))            //(g << 2): 16
	fmt.Printf("(h >> 2): %v\n", (h >> 2))            //(h >> 2): 2

	var i int
	i = 100
	fmt.Printf("i: %v\n", i) //i: 100
	i += 1
	fmt.Printf("i: %v\n", i) //i: 101
	i -= 1
	fmt.Printf("i: %v\n", i) //i: 100
	i *= 2
	fmt.Printf("i: %v\n", i) //i: 200
	i /= 2
	fmt.Printf("i: %v\n", i) //i: 100
}
