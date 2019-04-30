package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []user{{
	ID:   0,
	Name: "Alice",
}, {
	ID:   2,
	Name: "John",
}}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(users); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
