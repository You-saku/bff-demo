package main

import (
	"log"
	"net/http"

	"go-api/handlers"
)

func main() {
	/*
	* insert test data
	* INSERT INTO authors (id, name, bio) VALUES (1, 'author1', 'bio');
	* INSERT INTO authors (id, name, bio) VALUES (2, 'author2', 'bio');
	 */
	http.HandleFunc("/api/authors", handlers.ListAuthors)
	log.Fatal(http.ListenAndServe(":80", nil))
}
