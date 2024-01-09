package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "hi this is radha"

	words := strings.Fields(input)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(strings.Join(words," "))
}
