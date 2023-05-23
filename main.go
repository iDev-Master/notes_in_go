package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/behruzjon/go_notes/database"

	_ "github.com/lib/pq"
)

func main() {

	database.ConnectDB()

	// defer db.Close()

	// fs := http.FileServer(http.Dir("."))
	// http.Handle("/", fs)

	// Serve the index.html file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Start the server
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
