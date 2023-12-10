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
	count := 1
	n := &Node{
		data: count,
	}
	n.next = nil

	ll.head = n
	count = count+1
	curr := ll.head
	for count <= d {
		n = &Node{
		data: count,
	}
		n.next = nil
		curr.next = n

		curr = n
		count = count+1
		//fmt.Println(n.data)
	}

	curr.next = ll.head
	}

	func (ll *LinkedList) show() {
		curr := ll.head
		for curr.next != ll.head {
		fmt.Println(curr.data)
		curr = curr.next
		}
		fmt.Println(curr.data)
		curr = curr.next

		for curr.next != ll.head {
			fmt.Println(curr.data)
			curr = curr.next
		}
		fmt.Println(curr.data)
	}
	func (ll *LinkedList) survivor(n,k int) int {
		curr := ll.head
		for count:=n;count>1;count--{
		  for i:=0;i<k-1;i++ {
			curr = curr.next
		  }
		  fmt.Println(curr.data)
		  curr = curr.next.next

		}
		 fmt.Printf("the answer is %d", curr.data)
		 return curr.data
	   }


	func main()  {

	ll := &LinkedList{}
	ll.insert(7)

	ll.show()

	}