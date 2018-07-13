package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//create global database variable to reference in endpoint functions
var db *sql.DB

//API endpont which gets all blog posts within the database
func getBlogPosts(w http.ResponseWriter, req *http.Request) {
	rows, _ := db.Query("Select post_id, title, body FROM blog")
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
