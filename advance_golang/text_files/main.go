package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]int)
	file, err := os.Open("file1.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		m[word]++
	}
	for k,v := range m {
		fmt.Println(k,v)
	}

}
