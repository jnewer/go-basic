package main

import "fmt"

func main() {
	m1 := make(map[string]string, 0)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"
	fmt.Printf("m1: %v\n", m1) //m1: map[age:20 email:tom@gmail.com name:tom]

	name := m1["name"]
	age := m1["age"]
	email := m1["email"]
	fmt.Printf("name: %v\n", name)   //name: tom
	fmt.Printf("age: %v\n", age)     //age: 20
	fmt.Printf("email: %v\n", email) //email: tom@gmail.com

	m2 := map[string]string{
		"name":  "kite",
		"age":   "20",
		"email": "kite@gmai.com",
	}
	fmt.Printf("m2: %v\n", m2) //m2: map[age:20 email:kite@gmai.com name:kite]

	v, ok := m2["address"]
	if ok {
		fmt.Printf("键存在")
		fmt.Println(v)
	} else {
		fmt.Printf("键不存在")
	}

}
