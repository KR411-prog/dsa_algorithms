package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fs, err := os.Open("document")

	if err != nil {
		log.Fatal("Unable to open file")
	}
	scanner := bufio.NewScanner(fs)
	count := 0
	for scanner.Scan() {
		count++
	}

	fmt.Printf("Number of lines %d", count)

	fs,err = os.Create("document_2")
	if err != nil {
		log.Fatal("Unable to create file")
	}
	writer := bufio.NewWriter(fs)

	n,err := writer.WriteString("This is Radha. This is test")
	err = writer.Flush()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("number of bytes written into the file is %d", n)
}
