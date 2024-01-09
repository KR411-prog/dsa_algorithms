package main

import (
	"fmt"
	"sync"
)

func helloworld(wg *sync.WaitGroup) {

	defer wg.Done() //this will reduce the internal counter by 1
	fmt.Println("helloworld")
}

func goodbye(wg *sync.WaitGroup) {
	fmt.Println("goodbye")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg)
	go goodbye(&wg)
	//Note the internal counter
	wg.Wait() // This method blocks the execution of code until the internal counter becomes 0.
	//time.Sleep(1 * time.Second) //we can use Sleep which pauses the execution of the main function.
	//but we are actually blocking the execution of the program for 1 sec. At production level, each millisecond is vital.Millions of concurrent requests can create a huge problem.
	//thatswhy we should use sync.Waitgroup

}
