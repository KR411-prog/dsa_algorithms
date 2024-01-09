package main


func binarysearch(input []int, target int) int {
  if len(input) == 0 {
	return -1
  }
  if len(input) == 1 {
	return 0
  }
  low := 0
  high := (len(input)) - 1

  for low < high {
	mid := (low+high)/2
	if input[mid] == target {
		return mid
	}
	if input[mid] < target {
		low = mid+1
	} else if input[mid] > target {
		high = mid-1
	}

  }
  	if input[low] == target {
		return low
	} else {
		return -1
	}
}


