package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
)

var Client = ConnectRedis()

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return client
}

type Person struct {
	Id   int
	Name string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	var Student = struct {
		Id   int
		Name string
	}{
		1, "Shreyas",
	}
	json.NewEncoder(w).Encode(Student)
}

func Store(w http.ResponseWriter, r *http.Request) {
	var pers Person
	json.NewDecoder(r.Body).Decode(&pers)
	err := Client.Set(strconv.Itoa(pers.Id), pers.Name, 0).Err()
	if err != nil {
		log.Fatal("error setting value in redis", err)
	}
	json.NewEncoder(w).Encode("Value stored")
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	resp, err := Client.Get(id).Result()
	if err != nil {
		log.Fatal("error when getting the key", err)
	}
	idr, _ := strconv.Atoi(id)
	pers := Person{Id: idr, Name: resp}
	json.NewEncoder(w).Encode(pers)

}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/store", Store)
	http.HandleFunc("/get", Get)
	log.Println("Starting on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
