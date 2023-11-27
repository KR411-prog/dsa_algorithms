package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insert(d int) {
	n := &Node{data: d}
	n.next = nil

	if ll.head == nil {
		ll.head = n
		return
	}

	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = n
}

func (ll *LinkedList) delete(d int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == d {
		ll.head = ll.head.next
		return
	}
	curr := ll.head
	prev := curr

	for curr != nil && curr.data != d {
		prev = curr
		curr = curr.next
	}
	if curr.data == d {
		prev.next = curr.next
	}
}

func (ll *LinkedList) show() {
	if ll.head == nil {
		fmt.Println("")
		return
	}
	curr := ll.head
	for curr.next != nil {
		fmt.Printf("%d -->", curr.data)
		curr = curr.next
	}
	fmt.Printf("%d", curr.data)
}

func main() {
	ll := &LinkedList{}

	ll.insert(2)
	ll.insert(4)
	ll.show()
	fmt.Println("")
	ll.delete(2)
	ll.show()
	fmt.Println("")
	ll.insert(6)
	ll.show()
	fmt.Println("")
	ll.insert(8)
	ll.insert(10)
	ll.show()
}
