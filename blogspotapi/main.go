package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getBlogPosts(w http.ResponseWriter, req *http.Request) {

}

func createBlogPost(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/blgs", getBlogPosts).Methods("GET")
	router.HandleFunc("/blogs", createBlogPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
