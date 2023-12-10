package main

import (
	"fmt"
	"os"
)

func main() {
	s := "){"

	if len(s)%2 != 0 {
		fmt.Println("false")
		os.Exit(1)
	}

	m := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]rune,0)

	for _,c := range s {
		if c == '(' || c=='[' || c=='{' {
			stack = append(stack,c)
		} else if  c == ')' || c=='}' || c==']' {
			if len(stack) == 0 {
				fmt.Println(false)
				os.Exit(1)
			}
			if stack[len(stack)-1] == m[c] {
				stack = stack[:len(stack)-1]
			} else {
				fmt.Println(false)
				os.Exit(1)
			}
		}
		fmt.Println(stack)
	}
	if len(stack) != 0 {
		fmt.Println(false)
		os.Exit(1)
	}
	fmt.Println(true)
	fmt.Println(stack)

}
