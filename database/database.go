package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB // holding my db connection

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "01"
	dbname   = "super_notes"
)

// func ConnectDB() {

// 	// Connecting to postrgesql
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// Open the database connection
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Ping the database to verify the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create the database if it doesn't exist
// 	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS super_notes")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print a success message
// 	log.Println("Connected to the PostgreSQL database")
// }

func ConnectDB() {
	// Connect to the default "postgres" database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS (SELECT FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	// If the database doesn't exist, create it
	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created database:", dbname)
	}

	// Connect to the specified database
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the PostgreSQL database")
}
