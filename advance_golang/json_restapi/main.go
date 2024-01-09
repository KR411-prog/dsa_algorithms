package main

import (
	"fmt"
	"net/http"
)

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Bird struct {
// 	Species     string
// 	Description string
// }

// func main() {
// 	birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
// 	var birds []Bird
// 	json.Unmarshal([]byte(birdJson), &birds)
// 	fmt.Printf("Birds : %+v", birds)
// }

func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req,err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp,err := client.Do(req)

}
