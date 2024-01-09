package main

import (
	"fmt"
)

func main() {
	s := "test"

	for _, e := range s {
		if e == 't' {
			fmt.Println("Its t")
		}
		if e == rune('t') {
			fmt.Println("Its t rune")
		}

	}

	fmt.Println('t')
	fmt.Println(rune('t'))

}
