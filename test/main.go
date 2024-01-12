package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("document")
	if err != nil {
		log.Println("unable to open file")
	}

	f2, err := os.Create("newdocument")

	//fmt.Println(scanner.Text())
	if _, err := io.Copy(bufio.NewWriter(f2), bufio.NewReader(f)); err != nil {
		fmt.Errorf("cannot copy file from %q to %q: %w", "document","newdocument", err)
	}

	f2.Close()
	f.Close()

}
