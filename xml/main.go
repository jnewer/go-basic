package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `zml:"age"`
	Email   string   `xml:"email"`
}

func Marshal() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}

	b, _ := xml.MarshalIndent(p, " ", "  ")
	fmt.Printf("b: %v\n", string(b))
}

func read() {
	b, _ := ioutil.ReadFile("a.xml")
	var p Person
	xml.Unmarshal(b, &p)
	fmt.Printf("p: %v\n", p) //p: {{ person} tom 0 tom@gmail.com}
}

func write() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "tom@gmail.com",
	}
	f, _ := os.OpenFile("a.xml", os.O_WRONLY, 0777)
	defer f.Close()
	e := xml.NewEncoder(f)
	e.Encode(p)
}

func main() {
	// Marshal()
	// read()
	write()
}
