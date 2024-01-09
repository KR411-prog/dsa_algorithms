package main

import (
	"net/http"
	"encoding/json"
)

type Animal struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}

func AnimalHandler(w http.ResponseWriter, r *http.Request) {
	animals := []Animal{
		Animal{"Zebra", "Herbi"},
		Animal{"Lion", "carnivorous"},
	}

	json.NewEncoder(w).Encode(animals)

}

func main() {
	http.HandleFunc("/animals", AnimalHandler)
	http.ListenAndServe(":1234", nil)
}
