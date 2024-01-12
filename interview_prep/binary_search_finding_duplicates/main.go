
package main

func worst_case_algorithm_FindDuplicates(arr1 []int, arr2 []int) []int {
	//O(m+n)
	var i,j int
	result := []int{}
	for i < len(arr1) && j< len(arr2) {
	  if arr1[i] == arr2[j] {
		result = append(result,arr1[i])
		i++
		j++
	  } else if arr1[i] < arr2[j] {
		  i++
	  } else {
		j++
	  }

	}
	return result
  }

func search(arr []int, e int) bool {
//O(nlog(m)) m is len of longest array
  i := 0
  j := len(arr)-1
  var mid int
  for i<=j {
    mid = (i+j) /2

    if e == arr[mid] {
      return true
    } else if e < arr[mid] {
      j = mid-1
    } else {
      i = mid+1
    }

  }
  return false
}

func FindDuplicates(arr1 []int, arr2 []int) []int {
  result := []int{}
  if len(arr1) < len(arr2) {
    for _,e := range arr1 {
      if search(arr2, e) {
        result = append(result,e)
      }
    }
  } else {
    for _,e := range arr2 {
      if search(arr1, e) {
         result = append(result,e)
      }
    }
  }

  return result
}

func main() {

}