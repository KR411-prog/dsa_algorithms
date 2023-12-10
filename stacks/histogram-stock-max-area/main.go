package main

import "fmt"

func main() {
	//  stack := []int{}

	//  stack = append(stack,1)
	//  stack = append(stack,2)
	//  stack = append(stack,3)
	//  stack = append(stack,4)
    //  fmt.Println(stack)

	stack := make([]int,5)
	stack = append(stack,1)
	stack = append(stack,40)
	stack = append(stack,100)
	stack = append(stack,200)
	fmt.Println(stack)
}