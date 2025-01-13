package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
)

func main() {
	// Define the connection string
	connStr := "postgres://admin:admin@localhost:5432/brownie_points?sslmode=disable"

	// Open the database connection
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	// Configure the connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Run a SELECT query
	query := `SELECT id, username, password FROM "User";`
	rows, err := db.Query(query) // Example parameter: `active = true`
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	// Iterate through the rows
	fmt.Println("Active Users:")
	for rows.Next() {
		var id string
		var username string
		var password string
		if err := rows.Scan(&id, &username, &password); err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Printf("ID: %s, User Name: %s, Password: %s\n", id, username, password)
	}

	// Check for errors after iterating
	if err := rows.Err(); err != nil {
		log.Fatalf("Error occurred during iteration: %v", err)
	}
}
