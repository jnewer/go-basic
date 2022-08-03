package main

import "fmt"

type Person struct {
	id, age     int
	name, email string
}

func main() {
	var tom Person
	fmt.Printf("tom: %v\n", tom) //tom: {0 0  }
	kite := Person{}
	fmt.Printf("kite: %v\n", kite) //kite: {0 0  }

	tom.id = 1
	tom.name = "tom"
	tom.age = 20
	tom.email = "tom@gmail.com"
	fmt.Printf("tom: %v\n", tom) //tom: {1 20 tom tom@gmail.com}

	var dog struct {
		id   int
		name string
	}
	dog.id = 1
	dog.name = "dog"
	fmt.Printf("dog: %v\n", dog) //dog: {1 dog}
}
