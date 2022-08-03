package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat...")
}
func (a Animal) sleep() {
	fmt.Println("sleep...")
}

type Dog struct {
	Animal
}

type Cat struct {
	Animal
}

func main() {
	dog := Dog{
		Animal{
			name: "dog",
			age:  2,
		},
	}

	cat := Cat{
		Animal{
			name: "cat",
			age:  3,
		},
	}

	dog.eat()
	dog.sleep()

	cat.eat()
	cat.sleep()
}
