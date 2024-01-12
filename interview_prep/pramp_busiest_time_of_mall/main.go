package main

/*
input:  data = [ [1487799425, 14, 1],
                 [1487799425, 4,  0],
                 [1487799425, 2,  0],
                 [1487800378, 10, 1],
                 [1487801478, 18, 0],
                 [1487801478, 18, 1],
                 [1487901013, 1,  0],
                 [1487901211, 7,  1],
                 [1487901211, 7,  0] ]

output: 1487800378 # since the increase in the number of people
                   # in the mall is the highest at that point

*/
func FindBusiestPeriod(data [][]int) int {
	i:=0
	var current,result,max_people int
	for i < len(data) {

		if data[i][2] == 1 {
			current = current+data[i][1]
		} else {
			current = current-data[i][1]
		}
		if  i+1 == len(data) || data[i][0] != data[i+1][0]  {
			if max_people < current {
				result = data[i][0]
				max_people = current
			}
	}
	 i++

}


   return result
}

func main() {

}