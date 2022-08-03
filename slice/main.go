package main

import "fmt"

func main() {
	var names []string
	var numbers []int

	fmt.Printf("names: %v\n", names)     //names: []
	fmt.Printf("numbers: %v\n", numbers) //numbers: []
	fmt.Println(names == nil)            //true
	fmt.Println(numbers == nil)          //true

	var names2 = []string{"tom", "kite"}
	var numbers2 = []int{1, 2, 3}
	fmt.Printf("names2 len: %d cap: %d\n", len(names2), cap(names2))       //names2 len: 2 cap: 2
	fmt.Printf("numbers2 len: %d cap: %d\n", len(numbers2), cap(numbers2)) //numbers2 len: 3 cap: 3

	var s1 = make([]string, 2, 3)
	fmt.Printf("s1 len: %d cap: %d\n", len(s1), cap(s1)) //s1 len: 2 cap: 3
}
