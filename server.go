package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"learning/scrape/types"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	url := "https://www.reddit.com/.json"
	p := make(chan *types.ApiResponse)
	go getPosts(url, p)
	posts := json.NewEncoder(w).Encode(p)
	fmt.Fprintln(w, posts)
}

func server() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
