package main

import "fmt"

func main() {
	// var a [3]int    // 定义一个int类型的数组a，长度是3
	// var s [2]string // 定义一个字符串类型的数组s，长度是2
	// var b [2]bool

	// fmt.Printf("a: %T\n", a)
	// fmt.Printf("s: %T\n", s)
	// fmt.Printf("b: %v\n", b)

	// var a = [3]int{1, 2, 3}
	// var s = [2]string{"tom", "kite"}
	// var b = [2]bool{true, false}

	// a1 := [2]int{1, 2} // 类型推断

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("s: %v\n", s)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("a1: %v\n", a1)

	// var a = [...]int{1, 2, 3}
	// var s = [...]string{"tom", "kite"}
	// var b = [...]bool{true, false}

	// a1 := [...]int{1, 2} // 类型推断

	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("s: %v\n", s)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("a1: %v\n", a1)

	var a = [...]int{0: 1, 2: 2}
	var s = [...]string{1: "tom", 2: "kite"}
	var b = [...]bool{2: true, 5: false}

	a1 := [...]int{1, 2} // 类型推断

	fmt.Printf("a: %v\n", a)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a1: %v\n", a1)

}
