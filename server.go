package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func PostIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Posts Index")
}

func PostShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}

func server() {
	url := "https://www.reddit.com/.json"
	go getPosts(url)
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/posts/:id", PostIndex)
	log.Fatal(http.ListenAndServe(":8080", router))
}
