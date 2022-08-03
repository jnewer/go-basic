package main

import "fmt"

type Player interface {
	playMusic()
}
type Video interface {
	playVideo()
}
type Mobile struct {
}

func (m Mobile) playMusic() {
	fmt.Println("播放音乐")
}

func (m Mobile) playVideo() {
	fmt.Println("播放视频")
}

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()

	var p Pet
	p = Cat{}
	p.eat()
	p = Dog{}
	p.eat()
}
