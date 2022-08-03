package main

import (
	"fmt"
)

type Person struct {
	name string
}

func showPerson(person Person) {
	fmt.Printf("person: %p\n", &person)
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}
func showPerson2(person *Person) {
	fmt.Printf("person: %p\n", person)
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}

func (person Person) showPerson3() {
	fmt.Printf("person: %p\n", &person)
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}
func (person *Person) showPerson4() {
	fmt.Printf("person: %p\n", person)
	person.name = "kite"
	fmt.Printf("person: %v\n", person)
}

func main() {
	p1 := Person{name: "tom"}
	fmt.Printf("p1: %T\n", p1) //p1: main.Person
	p2 := &Person{name: "tom"}
	fmt.Printf("p2: %T\n", p2) //p2: *main.Person
	fmt.Println("---------------")
	p3 := Person{name: "tom"}
	fmt.Printf("p3: %p\n", &p3) //person: 0xc000050280
	showPerson(p3)              //person: 0xc000050280 person: {kite}
	fmt.Printf("p3: %v\n", p3)  //p3: {tom}
	fmt.Println("---------------")
	p4 := &Person{name: "tom"}
	fmt.Printf("p4: %p\n", p4) //p4: 0xc0000502b0
	showPerson2(p4)            //person: 0xc0000502b0 person: &{kite}
	fmt.Printf("p4: %v\n", p4) //p4: &{kite}
	fmt.Println("---------------")
	p5 := Person{name: "tom"}
	fmt.Printf("p5: %p\n", &p5) //p5: 0xc0000502e0
	p5.showPerson3()            //person: 0xc0000502f0 person: {kite}
	fmt.Printf("p5: %v\n", p5)  //p5: {tom}
	fmt.Println("---------------")
	p6 := &Person{name: "tom"}
	fmt.Printf("p6: %p\n", p6) //p6: 0xc000050320
	p6.showPerson4()           //person: 0xc000050320 person: &{kite}
	fmt.Printf("p6: %v\n", p6) //p6: &{kite}
}
