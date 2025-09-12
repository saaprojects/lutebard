package main

import (
	"context"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("backend.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("SUPABASE_DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)
}