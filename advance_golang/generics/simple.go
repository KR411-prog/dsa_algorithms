// package main

// import "fmt"

// // Parameterized Types
// type Radius interface {
// 	int64 | int8 | float64
// }

// func generic_circumference[R Radius](radius R) {
// 	var c R
// 	c = 2 * 3 * radius
// 	fmt.Println("The circumference is: ", c)
// }

// func main() {
// 	var r int64
// 	r = 10.0
// 	generic_circumference(r)
// }

package main

import "fmt"

func PrintString(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
func PrintInt(s []int) {
	for _, v := range s {
		fmt.Println(v)
	}
}
func Print[T any](s []T){
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	PrintString([]string{"This","is","Radha"})
	PrintInt([]int{1,3,5,6})
	Print([]string{"This","is","Radha"})
	Print([]int{1,3,5,6})
}
