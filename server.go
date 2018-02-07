package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	posts, _ := json.Marshal(Posts)
	w.Write(posts)
}

func server() {
	url := "https://www.reddit.com/.json"
	go getPosts(url)
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
