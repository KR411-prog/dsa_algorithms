package main

// this works for both file and input string
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func countParagraphs(reader io.Reader) int {

	scanner := bufio.NewScanner(reader)

	var inputLines []string
	for scanner.Scan() {
		text := scanner.Text()
		inputLines = append(inputLines, text)
	}
	inputText := strings.Join(inputLines, "\n")

	// Split the text based on empty lines
	paragraphs := strings.Split(inputText, "\n\n")

	// Count the number of paragraphs
	return len(paragraphs)
}

func main() {
	fs, err := os.Open("../document")
	if err != nil {
		log.Fatal("unable to open document")
	}

	fmt.Println(countParagraphs(fs))

	input := `This is first paragraph.

	This is secnd para.sdjnclksdlksmvlksmvlsmxvlsvv
	xvxvxvvx

	This is third para.xjchvnxlkvnklxvnmlkxv
	vnxjkvnxkvnlk
	xvnxk

	This is fourth para.xjchvnxlkvnklxvnmlkxv
	vnxjkvnxkvnlk
	xvnxk

	This is fifth para.xjchvnxlkvnklxvnmlkxv
	vnxjkvnxkvnlk
	xvnxk`

	fmt.Println(countParagraphs(strings.NewReader(input)))
}
