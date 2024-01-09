// Polymorphism - Allows different types to satisfy the same interfacr
package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Bark")
}

type Cat struct {}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

// speakAndPrint takes a Speaker interface and calls its Speak method
func speakAndPrint(speaker Speaker) {
	speaker.Speak()
}

func main() {
	d := Dog{}
	c := Cat{}
	speakAndPrint(d)
	speakAndPrint(c)

}


