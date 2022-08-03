package main

import (
	"fmt"
)

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}
func f1() {
	fmt.Printf("我没有返回值，只是进行一些计算")
}
func f2() (name string, age int) {
	name = "kang"
	age = 20
	return name, age
}
func f3() (name string, age int) {
	name = "kang"
	age = 20
	return
}

func f4() (name string, age int) {
	n := "kang"
	a := 20
	return n, a
}

func f5() (int, int) {
	return 1, 2
}
func main() {
	_, x := f5()
	fmt.Printf("x: %v\n", x)

}
