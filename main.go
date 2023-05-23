package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/behruzjon/go_notes/database"

	_ "github.com/lib/pq"
)

func main() {

	database.ConnectDB()

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "postgres://postgres:01@localhost:5432/super_notes?sslmode=disable")

	// db, err := sql.Open("postgres", "postgres://super_man:001@localhost:5432/name_of_my_db?sslmode=disable") // server: postgre_notes
	// db, err := sql.Open("postgres", "postgres://my_login:my_password@localhost:5432/name_of_my_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Serve static files (e.g., index.html)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Start the server
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// package main

// import (
// 	"net/http"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	// Create a new HTTP server.
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Serve the index.html file.
// 		http.ServeFile(w, r, "index.html")
// 	})

// 	// Listen on port 8080.
// 	http.ListenAndServe(":8080", nil)
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world!")
// }

// func about(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "This is the about page.")
// }

// func number(w http.ResponseWriter, r *http.Request) {
// 	// Get the number from the URL.
// 	number := r.URL.Path[1:]

// 	// Write the number to the page.
// 	fmt.Fprintf(w, "The number is %s.", number)
// }

// func main() {
// 	http.HandleFunc("/", index)
// 	http.HandleFunc("/about", about)
// 	http.HandleFunc("/number/", number)

// 	http.ListenAndServe(":8080", nil)
// }
