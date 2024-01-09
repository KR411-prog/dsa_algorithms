package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Article struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Views int    `json:"views"`
}

type Articles []Article
func (a Articles)Len() int{
	return len(a)
}

func (a Articles) Less(i,j int) bool{
	return a[i].Views < a[j].Views
}

func (a Articles) Swap(i,j int) {
	 a[i],a[j] = a[j],a[i]
}

func main() {

	response, _ := http.Get("http://localhost:9090/articles")
	resp, _ := ioutil.ReadAll(response.Body)

	var articles Articles

	json.Unmarshal(resp, &articles)
	fmt.Println(articles)

	sort.Sort(articles)
	fmt.Println(articles)
}
