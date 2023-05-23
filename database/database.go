// package main

package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB // holding my db connection

func ConnectDB() {
	var err error

	// Set up the PostgreSQL connection string
	// connStr := "postgres://username:password@localhost:5432/postgres?sslmode=disable"

	connStr := "user=postgres password=01 dbname=super_notes sslmode=disable"

	// Open the database connection
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	// Print a success message
	log.Println("Connected to the PostgreSQL database")
}

// func initDB() *sql.DB {
// 	connStr := "user=postgres password=123 dbname=postgres sslmode=disable"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err != nil {
// 		log.Fatal("Couldn't connect to database", err.Error())
// 	}

// 	return db
// }

// func StartDbConnection() {
// 	database = initDB()
// }

// func GetDBConn() *sql.DB {
// 	return database
// }

// func CloseDbConnection() error {
// 	if err := GetDBConn().Close(); err != nil {
// 		fmt.Errorf("error occurred on database connection closing: %s", err.Error())
// 	}
// 	return nil
// }
