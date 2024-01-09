// Embedding interfaces is a way to compose interfaces in Go, allowing you to create larger interfaces by combining smaller ones. 
// This promotes code reuse and flexibility in designing your types.
package main

import "fmt"

type Walker interface {
	Walk()
}

type Speaker interface {
	Speak()
}

type Animal interface {
	Walker
	Speaker
}

type Dog struct{}

func (d Dog) Walk() {
	fmt.Println("Dog is walking")
}

func (d Dog) Speak() {
    fmt.Println("Dog is barking")
}

func main() {
	d := Dog{}
	d.Walk()
	d.Speak()
}