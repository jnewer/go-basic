package main

import "fmt"

type USB interface {
	read()
	write()
}
type Computer struct {
}

type Mobile struct {
}

func (c Computer) read() {
	fmt.Println("computer read...")
}

func (c Computer) write() {
	fmt.Println("computer write...")
}

func (m Mobile) read() {
	fmt.Println("mobile read...")
}

func (m Mobile) write() {
	fmt.Println("mobile write...")
}

type OpenClose interface {
	open()
	close()
}
type Door struct {
}

func (d Door) open() {
	fmt.Println("open door...")
}
func main() {
	c := Computer{}
	m := Mobile{}

	c.read()
	c.write()

	m.read()
	m.write()

	// var oc OpenClose
	// oc = Door{}
	// cannot use Door{} (value of type Door) as type OpenClose in assignment:
	// Door does not implement OpenClose (missing close method)
}
