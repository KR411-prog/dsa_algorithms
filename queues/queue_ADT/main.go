// FIFO
// using slices
//
// Structs
package main

import "fmt"

type Queue struct {
	data []int
	size int
}

func (q *Queue) enqueue(d int) error{
	if len(q.data) == q.size  {
		return fmt.Errorf("%s", "Queue overflow")
	}
	q.data = append(q.data,d)
	return nil
}

func (q *Queue) dequeue() (int,error) {
	if len(q.data) == 0 {
		return 0,fmt.Errorf("%s", "Queue underflow")
	}
	d := q.data[0]
	q.data = q.data[1:]
	return d,nil
}

func main() {
	queue := &Queue{
		size: 5,
     }
	 queue.enqueue(1)
	 queue.enqueue(2)
	 queue.enqueue(3)
	 queue.enqueue(4)
	 queue.enqueue(5)
	 fmt.Println(queue.enqueue(6))
	 fmt.Println(queue.dequeue())
	 fmt.Println(queue.dequeue())
	 fmt.Println(queue.dequeue())

	 fmt.Println(queue.dequeue())
	 fmt.Println(queue.dequeue())
	 fmt.Println(queue.dequeue())
}

