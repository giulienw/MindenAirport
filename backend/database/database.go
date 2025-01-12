package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Database struct {
	*sql.DB
}

func CreateConnection() Database {
	db, err := sql.Open("godror", os.Getenv("CONNECTIONSTRING"))

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Successfully connected to Oracle Database!")

	return Database{db}
}

func CloseConnection(db *sql.DB) {
	defer db.Close()
}
