package main

import "fmt"

type Person struct {
	name string
}

func (person Person) eat() {
	fmt.Println(person.name + " eating ...")
}

func (person Person) sleep() {
	fmt.Println(person.name + " sleep...")
}
func main() {
	var person Person
	person.name = "tom"
	person.eat()   //tom eating ...
	person.sleep() //tom sleep...
}
