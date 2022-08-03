package main

import "fmt"

var a int = intVar()

func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init")
}
func intVar() int {
	fmt.Println("init var...")
	return 100
}
func main() {
	fmt.Println("main...")
}
