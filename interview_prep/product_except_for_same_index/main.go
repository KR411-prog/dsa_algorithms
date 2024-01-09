//Array of Array Products
// Given an array of integers arr, you’re asked to calculate for each index i the product of all integers except the integer at that index (i.e. except arr[i]). Implement a function arrayOfArrayProducts that takes an array of integers and returns an array of the products.

// Solve without using division and analyze your solution’s time and space complexities.

package main

import "fmt"

func ArrayOfArrayProducts(arr []int) []int {

  if len(arr) == 0 || len(arr)==1 {
    return []int{}
  }
  left := make([]int,len(arr))
  right := make([]int,len(arr))

  result := make([]int,len(arr))
  left[0] = 1
  right[len(arr)-1] = 1
  for i:=1;i<len(arr);i++ {
    left[i]=left[i-1]*arr[i-1]
  }

  for j:=len(arr)-2;j>=0;j--{
    right[j]=right[j+1]*arr[j+1]
  }

  for i:=0;i<len(arr);i++ {
    result[i] = left[i] * right[i]
  }
  return result

}

func main() {
  fmt.Println(ArrayOfArrayProducts([]int{2, 7, 3, 4}))
}

