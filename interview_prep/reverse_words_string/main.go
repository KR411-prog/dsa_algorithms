//Input: s = “geeks quiz practice code”
//Output: s = “code practice quiz geeks”
package main

import (
	"fmt"
	"strings"
)
func reverse_words(s string) string{
	var result []string
	result = strings.Split(s," ")

	for i,j:=0,len(result)-1; i<j; i,j=i+1,j-1 {
		result[i],result[j] = result[j],result[i]
	}
	return fmt.Sprintf(strings.Join(result," "))
}
func main() {
	input := "geeks quiz practice code"

	fmt.Println(reverse_words(input))
}


