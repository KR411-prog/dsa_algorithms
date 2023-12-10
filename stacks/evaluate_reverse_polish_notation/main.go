// Algorithm is in the book
import (
	"strings"
	"strconv"
	)
	func evalRPN(tokens []string) int {
		stack := []int{}

		for _,c := range tokens {
		  fmt.Printf("c value is %v ,", c)
		  result := 0
		  d,err := strconv.Atoi(c)

		  if err != nil {
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]

			switch c {
			  case "+":
				result = operand2+operand1
			  case "-":
				result = operand2-operand1
			  case "*":
				result = operand2*operand1
			  case "/":
				result = operand2/operand1
			}
			stack = append(stack[:len(stack)-2],result)
		  } else {
			stack = append(stack, d)
		  }
		}
		return stack[0]
	}