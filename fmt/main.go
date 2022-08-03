package main

import (
	"fmt"
)

func test1() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scan(&age)
	fmt.Printf("age: %v\n", age)
}
func test2() {
	var name string
	fmt.Println("请输入姓名：")
	fmt.Scanf("%s", &name)
	fmt.Printf("name: %v\n", name)
}

func test3() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("age: %v\n", age)
}
func main() {
	// test1()
	// test2()
	test3()
}
