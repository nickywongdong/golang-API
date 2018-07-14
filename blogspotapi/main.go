package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//define a structure which all blog posts will need to adhere to:
type Post struct {
	post_id string `json:"post_id"`
	title   string `json:"title"`
	body    string `json:"body"`
}

//create global database variable to reference in endpoint functions
var db *sql.DB

//API endpont which gets all blog posts within the database
func getBlogPosts(w http.ResponseWriter, req *http.Request) {
	//create array which holds results:
	var posts []Post

	//select all rows from database with query command and store
	rows, _ := db.Query("Select post_id, title, body FROM blog")

	//iterate through rows, and append each to posts array to be returned later
	var tempPost Post
	for rows.Next() {
		rows.Scan(&tempPost.post_id, &tempPost.title, &tempPost.body)
		posts = append(posts, tempPost)
	}

	json.NewEncoder(w).Encode(posts)
}

func createBlogPost(w http.ResponseWriter, req *http.Request) {

}

func main() {
	//create router
	router := mux.NewRouter()

	//open and use database blog.db file through sqlite3
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal()
	}

	//reference functions as endpoints
	router.HandleFunc("/posts", getBlogPosts).Methods("GET")
	router.HandleFunc("/post", createBlogPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
