package main

import (
	"fmt"
	"sync"
)

func getNumber_channels() <-chan int {
	ch := make(chan int)
	go func() {
		ch <- 5
	}()

	return ch
}

func getNumber() int {
	var i int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i = 5
		wg.Done()
	}()
	wg.Wait()
	return i
}

func main() {
	fmt.Println(getNumber())
	fmt.Println(<-getNumber_channels())
}
