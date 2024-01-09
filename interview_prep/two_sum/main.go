func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for id,val := range nums {
		d := target-val
		if _,ok := m[d];!ok {
			m[val] = id
		} else {
			return []int{id,m[d]}
		}
	}
	return []int{-1,-1}
}	