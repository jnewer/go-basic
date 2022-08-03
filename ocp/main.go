package main

import "fmt"

type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}
func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}
func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
	name string
}

func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}
	person := Person{}

	person.care(dog)
	person.care(cat)
}
