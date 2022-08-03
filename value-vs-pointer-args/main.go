package main

import "fmt"

type Pet interface {
}

type Dog struct {
	name string
}

func (dog Dog) eat() {
	fmt.Printf("dog: %p\n", &dog)
	fmt.Println("dog eat...")
	dog.name = "黑黑"
}

func (dog *Dog) eat2() {
	fmt.Printf("dog: %p\n", dog)
	fmt.Println("dog eat...")
	dog.name = "黑黑"
}
func main() {
	dog := Dog{name: "花花"}
	fmt.Printf("dog: %p\n", &dog) //dog: 0xc000050250

	// dog: 0xc000050260
	// dog eat...
	dog.eat()
	fmt.Printf("dog: %v\n", dog) //dog: {花花}

	dog2 := &Dog{name: "花花"}
	fmt.Printf("dog: %p\n", dog2) //dog: 0xc000050280
	// dog: 0xc000050290
	// dog eat...
	dog2.eat2()
	fmt.Printf("dog: %v\n", dog2) //dog: &{黑黑}

}
