package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func Marshal() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b)) //b: {"Name":"tom","Age":20,"Email":"tom@gmail.com"}
}

func Unmarshal() {
	b1 := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var m Person
	json.Unmarshal(b1, &m)
	fmt.Printf("m: %v\n", m) //m: {tom 20 tom@gmail.com}
}

func test3() {
	b := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com","Parents":["tome","kite"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Printf("f: %v\n", f) //f: map[Age:20 Email:tom@gmail.com Name:tom Parents:[tome kite]]
}

func test4() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}
	b, _ := json.Marshal(p)
	fmt.Printf("b: %v\n", string(b)) //b: {"Name":"tom","Age":20,"Email":"tom@gmail.com","Parent":["big tom","big kite"]}
}

func test5() {
	// dec := json.NewDecoder(os.Stdin)
	// a.json : {"Name":"tom","Age":20,"Email":"tom@gmail.com", "Parents":["tom", "kite"]}
	f, _ := os.Open("a.json")
	dec := json.NewDecoder(f)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("v: %v\n", v)
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

func test6() {
	type Person struct {
		Name   string
		Age    int
		Email  string
		Parent []string
	}

	p := Person{
		Name:   "tom",
		Age:    20,
		Email:  "tom@gmail.com",
		Parent: []string{"big tom", "big kite"},
	}
	f, _ := os.OpenFile("a.json", os.O_WRONLY, 0777)
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(p)
}
func main() {
	// Marshal()
	// Unmarshal()
	// test3()
	// test4()
	// test5()
	test6()

}
