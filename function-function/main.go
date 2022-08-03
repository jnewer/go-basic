package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("name: %v\n", name)
}
func f1(name string, f func(string)) {
	f(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}
func main() {
	f1("tom", sayHello) //name: tom

	add := cal("+")
	addr := add(1, 2)
	fmt.Printf("addr: %v\n", addr) //addr: 3

	fmt.Println("-----------")

	sub := cal("-")
	subr := sub(100, 50)
	fmt.Printf("subr: %v\n", subr) //subr: 50

}
