package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	var Student = struct {
		Id   int
		Name string
	}{
		1, "Shreyas",
	}
	json.NewEncoder(w).Encode(Student)

}

func main() {
	http.HandleFunc("/hello", Hello)
	log.Println("Starting on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
