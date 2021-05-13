package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Body string `json:"text"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	var message = Message{
		Body: name + " on " + time.Now().Format("2006.01.02 15:04:05"),
	}

	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
