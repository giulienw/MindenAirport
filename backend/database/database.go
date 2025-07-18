// Package database provides database connection management and data access layer
// for the MindenAirport application. It handles Oracle database connections
// and provides a wrapper around sql.DB for easier testing and abstraction.
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Database wraps sql.DB to provide additional functionality and easier testing.
// It serves as the main interface for all database operations in the application.
type Database struct {
	*sql.DB
}

// CreateConnection establishes a connection to the Oracle database using
// the connection string from the CONNECTIONSTRING environment variable.
// It performs a ping to verify connectivity and returns a Database instance.
//
// The connection string should be in Oracle format:
// user/password@host:port/service_name
//
// Returns:
//   - Database: A wrapped database connection ready for use
//
// Panics if connection fails or database is unreachable.
func CreateConnection() Database {
	db, err := sql.Open("godror", os.Getenv("CONNECTIONSTRING"))

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	} else {
		fmt.Println("Successfully connected to Oracle Database!")
	}

	return Database{db}
}

// CloseConnection properly closes the database connection.
// This should be called when the application shuts down to clean up resources.
//
// Parameters:
//   - db: The sql.DB connection to close
func CloseConnection(db *sql.DB) {
	defer db.Close()
}
