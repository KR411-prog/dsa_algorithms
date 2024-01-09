package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Views int    `json:"views"`
}

func views(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		{1, "article1", 100},
		{2, "article2", 50},
		{3, "article3", 300},
	}

	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(articles)

	if err != nil {
		serv_error := fmt.Sprintf("%s", err)
		json.NewEncoder(w).Encode(serv_error)
		w.WriteHeader(500)
	}



}

func main() {
	http.HandleFunc("/articles", views)
	http.ListenAndServe(":9090", nil)

}
