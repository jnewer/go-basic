package main

import "fmt"

func main() {
	var name string
	name = "tom"
	var p_name *string
	p_name = &name
	fmt.Printf("name: %v\n", name)       //name: tom
	fmt.Printf("p_name: %v\n", p_name)   //p_name: 0xc000050250
	fmt.Printf("*p_name: %v\n", *p_name) //*p_name: tom

	type Person struct {
		id   int
		name string
	}

	var tom = Person{1, "tom"}
	var p_person *Person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)             //tom: {1 tom}
	fmt.Printf("p_person: %p\n", p_person)   //p_person: 0xc000004078
	fmt.Printf("*p_person: %v\n", *p_person) //*p_person: {1 tom}

	var pperson = new(Person)
	fmt.Printf("pperson: %T\n", pperson)   //pperson: *main.Person
	fmt.Printf("pperson: %p\n", pperson)   //pperson: 0xc0000040c0
	fmt.Printf("*pperson: %v\n", *pperson) //*pperson: {0 }

	pperson.id = 1
	pperson.name = "tom"
	fmt.Printf("*pperson: %v\n", *pperson) //*pperson: {1 tom}

}
