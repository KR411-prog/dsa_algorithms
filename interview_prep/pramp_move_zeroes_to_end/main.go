package main

import "fmt"

func moveZeroes(arr []int) {
	nonZeroIndex := 0

	// Traverse the array and swap non-zero elements to the front
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			// Swap non-zero element with the element at nonZeroIndex
			arr[i], arr[nonZeroIndex] = arr[nonZeroIndex], arr[i]
			nonZeroIndex++
		}
	}
}

func main() {
	// Example array
	arr := []int{0, 2, 0, 4, 0, 6, 8, 0}

	fmt.Println("Original array:", arr)

	// Move zeroes to the end using swap
	moveZeroes(arr)

	fmt.Println("Array after moving zeroes to the end:", arr)
}