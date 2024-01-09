package main

import "fmt"

func reverse_string(s string) {
	r := []rune(s)
	//fmt.Printf("%T %v",r[0],string(r[0]))
	for i,j:=0,len(s)-1; i<=j; i,j=i+1,j-1 {

		r[i],r[j] = r[j],r[i]
	}
	fmt.Println(string(r))
}

func main() {
	s := "radha"
	reverse_string(s)
}
