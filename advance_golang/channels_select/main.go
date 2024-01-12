// https://www.golang-book.com/books/intro/10
package main

import (
	"fmt"
	"time"
)

func main() {
	// This creates a buffered channel with a capacity of 1. Normally channels are synchronous; both sides of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.
	c1 := make(chan string,1)
	c2 := make(chan string,1)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)

			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				time.Sleep(time.Second)
				fmt.Println("nothing ready")
			}
		}
	}()
	var input string 
	fmt.Scanln(&input)
}
