package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Typicode struct {
	UserId    int64  `json:"userId"`
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getStatusCode(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 500, fmt.Errorf("%s", "url is not working")
	}
	resp, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()
	//fmt.Println(string(resp))

	t := Typicode{}

	json.Unmarshal(resp, &t)
	return response.StatusCode, nil
}

func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatal("url not provided")
	// }
	url := flag.String("url", "https://jsonplaceholder.typicode.com/todos/2", "rest api url")
	flag.Parse()

	fmt.Printf("url is %s ", *url)

	resp, err := getStatusCode(*url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
