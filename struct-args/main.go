package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(person Person) {
	person.id = 1
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}

func showPerson2(person *Person) {
	person.id = 1
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}

func main() {
	person := Person{1, "tom"} //person: {1 tom}
	fmt.Printf("person: %v\n", person)
	fmt.Println("----------------")
	showPerson(person) //person: {1 kite}
	fmt.Println("----------------")
	fmt.Printf("person: %v\n", person) //person: {1 tom}
	fmt.Println("----------------")
	person2 := Person{1, "tom"} //person2: {1 tom}
	fmt.Printf("person2: %v\n", person2)
	fmt.Println("----------------")
	showPerson2(&person2) //person: &{1 kite}
	fmt.Println("----------------")
	fmt.Printf("person2: %v\n", person2) //person2: {1 kite}
}
