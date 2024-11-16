package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "Meles", Age: 25}
	jsonString, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonString)
}

func secondHandler(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "Meseret", Age: 13}
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(person.Name)

}

func decoderFunc(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(person.Name)
}
func main() {
	//Json encoding test and learn
	http.HandleFunc("/", handler)
	http.HandleFunc("/s", secondHandler)
	http.HandleFunc("/e", decoderFunc)

	http.ListenAndServe(":8080", nil)
}
