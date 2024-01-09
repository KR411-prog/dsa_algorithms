package main

import "fmt"
func indexEqualsValueSearch(arr []int) int{
  i := 0
  j := len(arr)-1
  for i<=j {
	mid := (i + j)/2
	if arr[mid] == mid {
		return mid
	}
	if arr[mid] > mid {
		j = mid-1
	}else {
		i = mid+1
	}
  }
  return -1
}

func main() {
  fmt.Println(indexEqualsValueSearch([]int{0,4,5,9}))
}