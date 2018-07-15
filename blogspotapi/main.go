package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//define a structure which all blog posts will need to adhere to:
type Post struct {
	Post_id int    `json:"post_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}

//create global database variable to reference in endpoint functions
var db *sql.DB

/*
 * Route to display all blog posts within the database
 */
func getBlogPosts(w http.ResponseWriter, req *http.Request) {
	//create array which holds results:
	var posts []Post

	//select all rows from database with query command and store
	var err error
	rows, err := db.Query("Select post_id, title, body FROM posts")
	if err != nil { //perhaps rows are empty
		log.Fatal("Error, could not retrieve any rows from database - ", err)
	}

	//iterate through rows, and append each to posts to an array to be returned later
	var tempPost Post
	for rows.Next() {
		rows.Scan(&tempPost.Post_id, &tempPost.Title, &tempPost.Body)
		posts = append(posts, tempPost)
	}

	rows.Close()

	json.NewEncoder(w).Encode(posts)
}

/*
 * Route to create a blog post
 */
func createBlogPost(w http.ResponseWriter, req *http.Request) {
	//create a post variable, and store body of http req in it
	var tempPost Post
	var err error

	err = json.NewDecoder(req.Body).Decode(&tempPost)

	if err != nil {
		log.Fatal("Error, could not save body of request - ", err)
	}

	//sql query to store new blog post
	statement, err := db.Prepare("INSERT INTO posts (title, body) VALUES (?, ?)")
	statement.Exec(tempPost.Title, tempPost.Body)

	if err != nil {
		log.Fatal("Error, could not insert post - ", err)
	} else {
		json.NewEncoder(w).Encode("Blog post successfully posted!")
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

	db.Close()
}
