package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Speak")
}

//inheritance
type Dog struct {
	Animal
	Id int
}

//Speak is Overridden by Dog Type
func (d Dog) Speak() {
	fmt.Println("Bark")
}

func (d Dog) Eat() {
	fmt.Println("Dog food")
}

func main() {
	var a Animal

	a.Name = "Cat"
	a.Speak()

	var d Dog
	d.Name = "Dog"
	d.Id = 1

	d.Speak()
	d.Eat()
}