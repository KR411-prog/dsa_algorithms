package main

/*
input:  slotsA = [[10, 50], [60, 120], [140, 210]]
        slotsB = [[0, 15], [60, 70]]
        dur = 8
output: [60, 68]

input:  slotsA = [[10, 50], [60, 120], [140, 210]]
        slotsB = [[0, 15], [60, 70]]
        dur = 12
output: [] # since there is no common slot whose duration is 12

*/

func min(x,y int) int {
  if x<y {
    return x
  }
  return y
}

func max(x,y int) int {
  if x>y {
    return x
  }
  return y
}

func MeetingPlanner(slotsA [][]int, slotsB [][]int, dur int) []int {
  var ia,ib int

  for ia<len(slotsA) && ib<len(slotsB) {
    start := max(slotsA[ia][0], slotsB[ib][0])
    end := min(slotsA[ia][1], slotsB[ib][1])

    if start+dur <= end {
      return []int{start,start+dur}
    }

    if slotsA[ia][1] < slotsB[ib][1] {
      ia++
    } else {
      ib++
    }

  }

  return []int{}



}

func main() {

}