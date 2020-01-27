package main

import (
	"encoding/json"
	"net/http"
)

// Post model.
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{
		Post{
			ID:    1,
			Title: "Title 1",
			Text:  "Text 1",
		},
	}
}

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "error": "Error marshalling the posts array" }`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
