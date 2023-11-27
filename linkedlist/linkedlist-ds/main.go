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
//nth node from end of linked list with two traversals
func (ll *LinkedList) nth_node_end(n int) (int,error) {
	var len int
	curr := ll.head

	for curr != nil {
		len++
		curr = curr.next
	}

	if len < n {
		return 0,fmt.Errorf("%s", "not enough nodes")
	}
	curr = ll.head
	count := 0
	for curr != nil && count < len-n {

		curr = curr.next
		count += 1

	}
	fmt.Println(count)
	if count == len-n {
		return curr.data,nil
	}
	return 0,nil
}

// nth node from end of linked list with one traversal - two pointer solution
func (ll *LinkedList) nth_node_end_two_ptr(n int) (int,error) {
  ptr := ll.head
  ptr_temp := ll.head
  count,len := 0,0
  curr := ll.head

  for curr != nil {
	  len++
	  curr = curr.next
  }

  if len < n {
	  return 0,fmt.Errorf("%s", "not enough nodes")
  }


  for count != n {
	ptr_temp = ptr_temp.next
	count += 1
  }

  for ptr_temp != nil {
	ptr = ptr.next
	ptr_temp = ptr_temp.next
  }
  return ptr.data,nil
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
	fmt.Println("")
	ll.insert(12)
	ll.insert(14)
	ll.show()
	fmt.Println("")
	d,err := ll.nth_node_end(7)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}


	d,err =  ll.nth_node_end_two_ptr(7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(d)
	}

}
