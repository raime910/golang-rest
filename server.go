package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8080"

	router := mux.NewRouter()
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Hello from the server.")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
