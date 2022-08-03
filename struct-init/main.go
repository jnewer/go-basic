package main

import "fmt"

func main() {
	type Person struct {
		id, age     int
		name, email string
	}

	var tom Person
	fmt.Printf("tom: %v\n", tom) //tom: {0 0  }

	kite := Person{
		id:    1,
		name:  "kite",
		age:   20,
		email: "kite@gmail.com",
	}

	fmt.Printf("kite: %v\n", kite) //kite: {1 20 kite kite@gmail.com}

	kite = Person{
		1,
		20,
		"kite",
		"kite@gmail.com",
	}

	fmt.Printf("kite: %v\n", kite) //kite: {1 20 kite kite@gmail.com}

	kite = Person{
		id:   1,
		name: "kite",
	}

	fmt.Printf("kite: %v\n", kite) //kite: {1 0 kite }
}
