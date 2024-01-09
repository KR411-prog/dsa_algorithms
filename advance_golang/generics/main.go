// https://www.freecodecamp.org/news/generics-in-golang/
package main

import (
	"fmt"
)

type Person interface {
	Work()
}

type worker string


func (w worker) Work(){
	fmt.Printf("%s is working\n", w)
}

func DoWork[T Person](things []T) {
    for _, v := range things {
        v.Work()
    }
}

//But even without generics it works. Any type that implements Person interface, can be passed to this function
func DoWorkInterface(things []Person) {
    for _, v := range things {
        v.Work()
    }
}


func main() {
	var a,b,c worker
	a = "A"
	b = "B"
	c = "C"
	DoWork([]worker{a,b,c})

	// var d,e,f Person
	// d=worker("D")
	// e=worker("E")
	// f=worker("F")

	var d,e,f worker
	d="D"
	e="E"
	f="F"
	DoWorkInterface([]Person{d,e,f})
}