package main

import "fmt"

func main() {
	//append
	// s1 := []int{1, 2, 3}
	// i := append(s1, 4)
	// fmt.Printf("i: %v\n", i) //i: [1 2 3 4]

	// s2 := []int{7, 8, 9}
	// i2 := append(s1, s2...)
	// fmt.Printf("i2: %v\n", i2) //i2: [1 2 3 7 8 9]

	//len 返回，数组、切片、字符串、通道的长度
	// s1 := "hello world"
	// i := len(s1)
	// fmt.Printf("i: %v\n", i) //i: 11

	// s2 := []int{1, 2, 3}
	// fmt.Printf("len(s2): %v\n", len(s2)) //len(s2): 3

	//print println
	// name := "tom"
	// age := 20
	// print(name, " ", age, "\n")
	// fmt.Println("---------")
	// println(name, " ", age)

	// // panic
	// defer fmt.Println("panic 异常后执行...")
	// panic("panic 错误...")
	// fmt.Println("end..")

	b := new(bool)
	fmt.Println(*b) //false

	i := new(int)
	fmt.Println(*i) //0
	s := new(string)
	fmt.Println(*s)

	var p *[]int = new([]int)
	fmt.Printf("p: %v\n", p) //p: &[]
	var v []int = make([]int, 10)
	fmt.Printf("v: %v\n", v) //v: [0 0 0 0 0 0 0 0 0 0]

}
