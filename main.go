package main

import "fmt"

type Speakable interface {
	Speak()
}

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println(c.Name + ": Meow")
}

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name + ": Howl")
}

func main() {
	var cat Speakable
	cat = Cat{ Name: "Gigy" }
	fmt.Println("Hello, <Main.go>")
	cat.Speak()
}