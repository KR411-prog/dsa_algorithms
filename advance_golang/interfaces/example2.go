package main

// https://www.alexedwards.net/blog/interfaces-explained

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// But the key thing to take away is that by using a interface type in our WriteLog() function declaration,
// we have made the function agnostic (or flexible) about the exact type of object it receives. All that matters is what methods it has.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

func example2() {
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	count := Count(3)
	WriteLog(count)
}
