package main

import "fmt"

func main() {
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"

	for key := range m {
		fmt.Println(key)
	}

	for key, value := range m {
		fmt.Println(key + ":" + value)
	}
}
