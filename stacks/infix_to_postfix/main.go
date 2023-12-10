// Infix to Postfix Algorithm
// Refer book
// 1. If any operand, dont push to stack but output
// 2. If any operator, push to stack. But if there is a higher precedence one, then pop it.
// Example: If you need to push +, but * is there in stack, then pop *. But not vice versa. Because * is higher precendence than +
// 3. If there is a (, push it and pop everything when you find a ")"

package main

import (
	"fmt"
)

func main() {
	input := "3*3/(7+1)"
	// 33
	stack := []string{}
    result := []string{}
	for _, c := range input {
		// fmt.Printf(" %s", string(c))
		// fmt.Println("")
		fmt.Print("\nstack is\n")
		fmt.Print(stack)
		if c >= 48 && c <= 57 {
			result = append(result,string(c))
			fmt.Println("result is")
			fmt.Print(result)
			continue
		}
		if string(c) == "(" {
			stack = append(stack, string(c))
		}
		if string(c) == "+" || string(c) == "-" {
			fmt.Println("inside + -")
			for len(stack) != 0 {
				fmt.Println("stack is")
				fmt.Print(stack)
				fmt.Println("")
				if stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/" || stack[len(stack)-1] == "^" || stack[len(stack)-1] == ")" {
					d := stack[len(stack)-1]
					result = append(result,d)
					fmt.Println("+ - result is")
					fmt.Print(result)
					stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(c))
			}
		}
		}
		if string(c) == "*" || string(c) == "/" || string(c) == "^" {

			for  len(stack) != 0 {
			 fmt.Printf("\nI am inserting %s\n", string(c))
			 if stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/" || stack[len(stack)-1] == "^" || stack[len(stack)-1] == ")" {
				d := stack[len(stack)-1]
				result = append(result,d)
				fmt.Println("result is")
				fmt.Print(result)
				stack = stack[:len(stack)-1]

			}

		}
		stack = append(stack, string(c))
	}
	}
	// count := len(stack) - 1
	// for count >= 0 {
	// 	fmt.Print(stack[count])
	// }


}