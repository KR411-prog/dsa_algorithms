package main

import (
	"fmt"
)

type Node struct {
	c    rune
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(d rune) {
	n := &Node{
		c: d,
	}
	n.next = s.top
	s.top = n
}

func (s *Stack) pop() (rune, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("%s", "stack is empty")
	}
	data := s.top.c
	s.top = nil
	return data, nil
}

func (s *Stack) isEmpty() bool {
	//fmt.Printf(" s.top value is , %v", s.top == nil)
	return s.top == nil
}
func main() {
	st := &Stack{}
	s := "()[[]]{4#}"
	ans := true
	for _, c := range s {

		if string(c) == "(" || string(c) == "[" || string(c) == "{" {
			fmt.Printf("\npushing %v", string(c))
			st.push(c)
		}

		if string(c) == ")" {
			result, err := st.pop()
			fmt.Println("")
			fmt.Println(string(result))
			fmt.Println(err)
			if err != nil || result != '(' {
				ans = false
				break
			}
		}

		if string(c) == "]" {
			result, err := st.pop()
			fmt.Println("")
			fmt.Println(string(result))
			fmt.Println(err)
			if err != nil || result != '['{
				ans = false
				break
			}
		}

		if string(c) == "}" {
			fmt.Println("")
			fmt.Println("I see a }")
			result, err := st.pop()
			fmt.Println("")
			fmt.Println(string(result))
			fmt.Println(err)
			if err != nil || result != '{' {
				ans = false
				break
			}
		}

	}
	//fmt.Println("hi")
	//fmt.Printf(" top value is end end %v", st.top == nil)

	if ans == false || !st.isEmpty() {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

}
