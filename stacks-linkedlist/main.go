//This Program is to implement stacks using linked list
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stack struct {
	head *Node
}

func (s *stack) push(d int) {
	n := &Node{data: d}

	if s.isEmpty() {
		n.next = nil
		s.head = n
		return
	}
	n.next = s.head
	s.head = n
}

func (s *stack) pop() (int,error) {
	if s.isEmpty() {
		return 0,fmt.Errorf("%s", "stack is empty")
	}
	d := s.head.data
	s.head = s.head.next
	return d,nil
}

func (s *stack) top() (int,error) {
	if s.isEmpty() {
		return 0,fmt.Errorf("%s", "stack is empty")
	}
	return s.head.data,nil
}

func (s *stack) isEmpty() bool {
	return s.head == nil
}
func main() {
	s := &stack{}
	s.push(2)
	s.push(4)
	fmt.Println(s.top())
	fmt.Println(s.pop())
	s.push(6)
	s.push(8)
	fmt.Println(s.pop())
	fmt.Println(s.top())
	fmt.Println(s.pop())
	fmt.Println(s.top())

	d,err:= s.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	d,err = s.top()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	d,err = s.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	d,err = s.top()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	d,err = s.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}
	s.push(10)
	s.push(12)
	fmt.Println(s.top())
}
