package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
	count int
}

func (q *Queue) enqueue(d int) {

	n := &Node{
		data: d,
	}

	if q.front == nil {
		q.front = n
		return
	}
	q.rear.next = n
	q.count++
	q.rear = n
}

func (q *Queue) dequeue() int {
	if q.front == nil {
		fmt.Println("empty queue")
		return -1
	}
	d := q.front.data
	q.front = q.front.next
	q.count--

	if q.front.next == nil {
		q.rear = nil
	}
	return d
}

func main() {

}
