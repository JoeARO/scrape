package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"learning/scrape/types"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	url := "https://www.reddit.com/.json"
	p := make(chan *types.ApiResponse)
	go getPosts(url, p)
	json.NewEncoder(w).Encode(p)
}

func server() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
