package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
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
	var err error
	rows, err := db.Query("Select post_id, title, body FROM posts")

	//iterate through rows, and append each to posts array to be returned later
	var tempPost Post

	//perhaps rows are empty
	if err != nil {
		log.Fatal("Error, could not retrieve any rows from database - ", err)
	}

	for rows.Next() {
		rows.Scan(&tempPost.post_id, &tempPost.title, &tempPost.body)
		fmt.Printf("%s", tempPost.post_id)
		fmt.Printf("%s", tempPost.title)
		fmt.Printf("%s", tempPost.body)

		posts = append(posts, tempPost)
	}

	json.NewEncoder(w).Encode(posts)
}

func createBlogPost(w http.ResponseWriter, req *http.Request) {
	//create a post variable, and store body of http req in it
	var tempPost Post
	_ = json.NewDecoder(req.Body).Decode(&tempPost)

	//sql query to store new blog post
	statement, err := db.Prepare("INSERT INTO posts (title, body) VALUES (?, ?)")
	statement.Exec(tempPost.title, tempPost.body)

	if err != nil {
		log.Fatal("Error, could not insert post - ", err)
	}
}

func main() {
	//create router
	router := mux.NewRouter()

	//open and use database blog.db file through sqlite3
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")

	if err != nil {
		log.Fatal("Error, could not open database - ", err)
	}

	//reference functions as endpoints
	router.HandleFunc("/posts", getBlogPosts).Methods("GET")
	router.HandleFunc("/post", createBlogPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
