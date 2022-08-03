package main

import (
	"fmt"
	"os"
)

func main() {
	//获取所有环境变量
	// s := os.Environ()
	// fmt.Printf("s: %v\n", s)

	s2 := os.Getenv("GOPATH")
	fmt.Printf("s2: %v\n", s2)

	os.Setenv("env1", "env1")
	s2 = os.Getenv("env1")
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("-----------")

	s3, b := os.LookupEnv("env1")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("s3: %v\n", s3)

	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	os.Unsetenv("env1")
	os.Unsetenv("NAME")
	os.Unsetenv("BURROW")
}
