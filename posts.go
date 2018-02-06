package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"learning/webscraper/types"
)

func getPosts(url string) *types.ApiResponse {
	for {
		// Create client so I can set User Agent for reddit API
		client := &http.Client{}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

		// Get posts in JSON format
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		// Extract JSON
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(fmt.Errorf("Body is nil"))
		}

		// Get Posts from the JSON body
		posts, err := extractPosts(body)
		if err != nil {
			panic(fmt.Errorf("No Posts!"))
		}
		fmt.Printf("%+v\n", posts)
		return posts
		time.Sleep(time.minute * 5)
	}
}

func extractPosts(body []byte) (*types.ApiResponse, error) {
	// Initialize ApiResponse Struct
	var posts = &types.ApiResponse{}
	// Unmarshal bytes to return posts
	err := json.Unmarshal(body, posts)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return posts, err
}
